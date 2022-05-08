package services

import (
	"blive/src/services/music"
	"blive/src/utils"
	"fmt"
	"github.com/botplayerneo/bili-live-api/log"
	"github.com/gobuffalo/packr/v2"
	"github.com/hajimehoshi/go-mp3"
	"github.com/q191201771/lal/pkg/httpflv"
	"github.com/q191201771/lal/pkg/remux"
	"github.com/q191201771/lal/pkg/rtmp"
	"os"
	"os/exec"
	"path"
	"runtime"
)

var CurrentVideoService *VideoService

func getMp3Duration(filename string) int {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		return 0
	}

	defer file.Close()
	d, err := mp3.NewDecoder(file)
	if err != nil {
		log.Error(err)
		return 0
	}
	return int(d.Length() / 4 / int64(d.SampleRate()))
}

var ch chan string
var DefaultMusic string = "354352"

func StartEncode(encodeChannel *chan music.SongDetail) {
	// 创建临时文件夹
	ch = make(chan string, 1000)
	encodePath := path.Join(os.TempDir(), "blive_encode")
	_ = os.MkdirAll(encodePath, os.ModePerm)
	log.Info(encodePath)

	// 保存资源文件到临时文件夹中
	resource := packr.New("Resource Box", "../resources")
	defaultAss, _ := resource.Find("default.ass")
	defaultAssPath := path.Join(encodePath, "default.ass")
	_ = utils.SaveBytesToFile(defaultAss, defaultAssPath)

	defaultJpg, _ := resource.Find("default.jpg")
	defaultJpgPath := path.Join(encodePath, "default.jpg")
	_ = utils.SaveBytesToFile(defaultJpg, defaultJpgPath)

	for {
		song := <-*encodeChannel
		songPath := path.Join(encodePath, song.Id)
		log.Infof("开始渲染歌曲 Id: %s Name: %s", song.Id, song.Name)

		// 将渲染需要的文件复制到临时文件夹
		_ = os.MkdirAll(songPath, os.ModePerm)
		_ = utils.CopyFile(song.LocalPath, path.Join(songPath, "1.mp3"))
		_ = utils.CopyFile(defaultJpgPath, path.Join(songPath, "1.jpg"))
		_ = utils.CopyFile(defaultAssPath, path.Join(songPath, "1.ass"))

		outPath := path.Join(songPath, "output.flv")
		if _, err := os.Stat(outPath); err == nil {
			log.Infof("歌曲存在跳过渲染 Id: %s Name: %s", song.Id, song.Name)
			ch <- outPath
			continue
		}

		// 构建渲染命令
		file, err := os.OpenFile(path.Join(songPath, "command.bat"), os.O_CREATE, 0777)
		if err != nil {
			log.Error(err)
			continue
		}

		_ = file.Truncate(0)
		musicDuration := getMp3Duration(song.LocalPath)

		_, _ = file.WriteString(
			fmt.Sprintf(
				"ffmpeg -thread_queue_size 24 -loop 1 -r 24 -t %d -f image2 -i 1.jpg -i 1.mp3 -vf ass=\"1.ass\" -s 1280x720 -pix_fmt yuvj422p -crf 28 -preset ultrafast -maxrate 2000k -bufsize 400000 -acodec aac -b:a 128k -c:v h264 -f flv output.flv -y",
				musicDuration,
			),
		)
		_ = file.Close()

		// 执行渲染命令
		var command *exec.Cmd
		if //goland:noinspection GoBoolExpressions
		runtime.GOOS == "windows" {
			command = exec.Command("cmd", "/c", "command.bat")
		} else {
			command = exec.Command("sh", "-c", "command.bat")
		}

		command.Dir = songPath
		//command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		_ = command.Start()
		command.Wait()

		// 通知给直播通道
		log.Infof("歌曲渲染完毕 Id: %s Name: %s Path: %s", song.Id, song.Name, outPath)
		ch <- outPath
	}
}

type VideoService struct {
	ps *rtmp.PushSession
}

func NewVideoService() VideoService {
	return VideoService{
		rtmp.NewPushSession(func(option *rtmp.PushSessionOption) {
			option.PushTimeoutMs = 0
			option.WriteAvTimeoutMs = 0
			option.WriteBufSize = 0
			option.WriteChanSize = 0
		}),
	}
}

func (v *VideoService) StartLiveStream() {
	err := v.ps.Push("rtmp://live-push.bilivideo.com/live-bvc/?streamname=live_927293_332_c521e483&key=71e3bf02c3502e40750dec010e364bee&schedule=rtmp&pflag=1")
	if err != nil {
		return
	}

	go func() {
		for {
			url := <-ch
			fmt.Println(url)
			flvFilePump := httpflv.NewFlvFilePump(func(option *httpflv.FlvFilePumpOption) {
				option.IsRecursive = false
			})

			tags, _ := httpflv.ReadAllTagsFromFlvFile(url)
			flvFilePump.PumpWithTags(tags, func(tag httpflv.Tag) bool {
				chunks := remux.FlvTag2RtmpChunks(tag)
				err := v.ps.Write(chunks)
				if err != nil {
					return false
				}
				return true
			})
		}
	}()

	_ = <-v.ps.WaitChan()
}
func (v *VideoService) StopLiveStream() {
	v.ps.Dispose()
}
