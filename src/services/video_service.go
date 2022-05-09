package services

//
//import (
//	"blive/src/database"
//	"blive/src/entities"
//	"blive/src/globals"
//	"blive/src/services/music"
//	"blive/src/utils"
//	"encoding/json"
//	"fmt"
//	"github.com/botplayerneo/bili-live-api/log"
//	"github.com/gobuffalo/packr/v2"
//	"github.com/hajimehoshi/go-mp3"
//	"github.com/q191201771/lal/pkg/httpflv"
//	"github.com/q191201771/lal/pkg/remux"
//	"github.com/q191201771/lal/pkg/rtmp"
//	"os"
//	"os/exec"
//	"path"
//	"runtime"
//)
//
//var ch chan string
//var DefaultMusic = "1937717747"
//var defaultMusicVideoPath string
//
//func getMp3Duration(filename string) int {
//	file, err := os.OpenFile(filename, os.O_RDONLY, 0777)
//	if err != nil {
//		return 0
//	}
//
//	defer file.Close()
//	d, err := mp3.NewDecoder(file)
//	if err != nil {
//		log.Error(err)
//		return 0
//	}
//	return int(d.Length() / 4 / int64(d.SampleRate()))
//}
//
//func StartEncode(encodeChannel *chan music.SongDetail) {
//	// 创建临时文件夹
//	ch = make(chan string, 1000)
//	homeDir, _ := os.UserHomeDir()
//	encodePath := path.Join(homeDir, "blive_tmp/blive_encode")
//	_ = os.MkdirAll(encodePath, os.ModePerm)
//	log.Info(encodePath)
//
//	// 保存资源文件到临时文件夹中
//	resource := packr.New("Resource Box", "../resources")
//	defaultAss, _ := resource.Find("default.ass")
//	defaultAssPath := path.Join(encodePath, "default.ass")
//	_ = utils.SaveBytesToFile(defaultAss, defaultAssPath)
//
//	defaultJpg, _ := resource.Find("default.jpg")
//	defaultJpgPath := path.Join(encodePath, "default.jpg")
//	_ = utils.SaveBytesToFile(defaultJpg, defaultJpgPath)
//
//	for {
//		song := <-*encodeChannel
//		songPath := path.Join(encodePath, song.Id)
//		log.Infof("开始渲染歌曲 Id: %s Name: %s", song.Id, song.Name)
//
//		// 将渲染需要的文件复制到临时文件夹
//		_ = os.MkdirAll(songPath, os.ModePerm)
//		_ = utils.CopyFile(song.LocalPath, path.Join(songPath, "1.mp3"))
//		_ = utils.CopyFile(defaultJpgPath, path.Join(songPath, "1.jpg"))
//		_ = utils.CopyFile(defaultAssPath, path.Join(songPath, "1.ass"))
//
//		outPath := path.Join(songPath, "output.flv")
//		if _, err := os.Stat(outPath); err == nil {
//			if song.Id == DefaultMusic {
//				defaultMusicVideoPath = outPath
//			}
//
//			log.Infof("歌曲存在跳过渲染 Id: %s Name: %s", song.Id, song.Name)
//			ch <- outPath
//			continue
//		}
//
//		// 构建渲染命令
//		shellName := "command.bat"
//		if runtime.GOOS != "windows" {
//			shellName = "command.sh"
//		}
//		file, err := os.OpenFile(path.Join(songPath, shellName), os.O_CREATE|os.O_RDWR, 0777)
//		if err != nil {
//			log.Error(err)
//			continue
//		}
//
//		_ = file.Truncate(0)
//		musicDuration := getMp3Duration(song.LocalPath) + 3
//
//		_, err = file.WriteString(
//			fmt.Sprintf(
//				"ffmpeg -thread_queue_size 256 -loop 1 -r 21 -t %d -f image2 -i 1.jpg -i 1.mp3 -vf ass=\"1.ass\" -s 1280x720 -pix_fmt yuvj422p -crf 28 -preset ultrafast -maxrate 2000k -bufsize 400000 -acodec aac -af \"apad=pad_dur=3\" -b:a 128k -c:v h264 -f flv output.flv -y",
//				musicDuration,
//			),
//		)
//		if err != nil {
//			log.Error(err)
//			continue
//		}
//		_ = file.Close()
//
//		// 执行渲染命令
//		var command *exec.Cmd
//		if //goland:noinspection GoBoolExpressions
//		runtime.GOOS == "windows" {
//			command = exec.Command("cmd", "/c", shellName)
//		} else {
//			command = exec.Command("sh", shellName)
//		}
//
//		command.Dir = songPath
//		//command.Stdout = os.Stdout
//		command.Stderr = os.Stderr
//		_ = command.Start()
//		command.Wait()
//
//		// 通知给直播通道
//		log.Infof("歌曲渲染完毕 Id: %s Name: %s Path: %s", song.Id, song.Name, outPath)
//		if song.Id == DefaultMusic {
//			defaultMusicVideoPath = outPath
//		}
//		ch <- outPath
//	}
//}
//
//type VideoService struct {
//	ps *rtmp.PushSession
//}
//
//func NewVideoService() VideoService {
//	return VideoService{nil}
//}
//
//func (v *VideoService) StartLiveStream() {
//	v.ps = rtmp.NewPushSession(func(option *rtmp.PushSessionOption) {
//		option.PushTimeoutMs = 10000
//		option.WriteAvTimeoutMs = 0
//		option.WriteBufSize = 2000
//		option.WriteChanSize = 128
//	})
//
//	var configs []database.SysConfig
//	var rtmpUrl string
//	database.DB.Where("name = 'live_room'").Find(&configs)
//	if len(configs) > 0 {
//		liveRoom := entities.LiveRoom{}
//		_ = json.Unmarshal([]byte(configs[0].Value), &liveRoom)
//		rtmpUrl = fmt.Sprintf("%s%s", liveRoom.LiveUrl, liveRoom.LivePassword)
//	}
//	if len(configs) == 0 || rtmpUrl == "" {
//		log.Error("无直播间配置，无法启动直播")
//		return
//	}
//
//	err := v.ps.Push(rtmpUrl)
//	if err != nil {
//		return
//	}
//
//	flvFilePump := utils.NewFlvFilePump(func(option *utils.FlvFilePumpOption) {
//		option.IsRecursive = false
//	})
//	defaultTags, _ := httpflv.ReadAllTagsFromFlvFile(defaultMusicVideoPath)
//	var currentTags []*[]httpflv.Tag
//
//	go func() {
//		for {
//			url := <-ch
//			fmt.Println(url)
//			tags, _ := httpflv.ReadAllTagsFromFlvFile(url)
//			currentTags = append([]*[]httpflv.Tag{&tags}, currentTags...)
//		}
//	}()
//
//	go func() {
//		flvFilePump.PumpWithTags(func(tag httpflv.Tag) bool {
//			chunks := remux.FlvTag2RtmpChunks(tag)
//			err := v.ps.Write(chunks)
//			if err != nil {
//				return false
//			}
//
//			if !globals.DanmuService.IsListening() {
//				return true
//			}
//			return true
//		}, func() *[]httpflv.Tag {
//			lens := len(currentTags)
//			if lens == 0 {
//				return &defaultTags
//			}
//
//			go func() {
//				currentTags = currentTags[:lens-1]
//			}()
//			return currentTags[lens-1]
//		})
//	}()
//
//	_ = <-v.ps.WaitChan()
//}
//func (v *VideoService) StopLiveStream() {
//	if v.ps != nil {
//		_ = v.ps.Dispose()
//	}
//}
