package services

import (
	"encoding/json"
	"fmt"
	"github.com/botplayerneo/bili-live-api/log"
	"github.com/imroc/req/v3"
	"os"
	"regexp"
	"strconv"
)

type MusicService struct {
	musicChannel *chan string
}

func NewMusicService(musicChannel *chan string) MusicService {
	service := MusicService{musicChannel}
	return service
}

func (m *MusicService) Start() {
	wg.Add(1)
	go func() {
		for {
			music, ok := <-*m.musicChannel
			if !ok {
				fmt.Printf("test: false")
				break
			}

			fmt.Printf("收到音乐: %s\n", music)
			isNumber, _ := regexp.MatchString("\\d+", music)
			if isNumber {
				musicId, _ := strconv.Atoi(music)
				GetMusic(musicId, nil)
			} else {
				searchResult := SearchMusic(music)
				if searchResult.Result.SongCount > 0 {
					firstSong := searchResult.Result.Songs[0]
					GetMusic(firstSong.Id, &firstSong)
				}
			}
		}

	}()
	wg.Wait()
}

type (
	NeteaseMusicData[T interface{}] struct {
		Result T   `json:"result"`
		Data   T   `json:"data"`
		Code   int `json:"code"`
	}
	NeteaseMusicSong struct {
		Id  int    `json:"id"`
		Url string `json:"url"`
	}
	NeteaseMusicSearchSong struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	NeteaseMusicSearch struct {
		SongCount int                      `json:"songCount"`
		Songs     []NeteaseMusicSearchSong `json:"songs"`
	}
)

func SearchMusic(keyword string) *NeteaseMusicData[NeteaseMusicSearch] {
	result := &NeteaseMusicData[NeteaseMusicSearch]{}
	resp, err := req.R().SetResult(&result).Get(fmt.Sprintf("https://netease-cloud-music-api-ochre-one.vercel.app/search?keywords=%s", keyword))
	if err != nil && resp != nil {
		return nil
	}

	listJson, _ := json.Marshal(result.Result)
	log.Infof("搜索到音乐信息 keyword: %s 列表信息: %s", keyword, listJson)
	return result
}

func GetMusic(id int, searchSong *NeteaseMusicSearchSong) {
	// 如果没有查询过音乐 则先查询
	if searchSong == nil {
		searchResult := SearchMusic(strconv.Itoa(id))
		if searchResult.Result.SongCount > 0 {
			searchSong = &searchResult.Result.Songs[0]
		}
	}

	// 通过ID查询音乐链接
	result := &NeteaseMusicData[[]NeteaseMusicSong]{}
	_, err := req.R().SetResult(&result).Get(fmt.Sprintf("https://netease-cloud-music-api-ochre-one.vercel.app/song/url?id=%d&br=320000", id))
	if err != nil {
		return
	}

	listJson, _ := json.Marshal(result.Data)
	log.Infof("获取到音乐信息 Id: %d 列表信息: %s", id, listJson)
	if len(result.Data) == 0 {
		return
	}

	music := result.Data[0]
	fileName := fmt.Sprintf("/music/%d.mp3", id)

	log.Infof("开始下载音乐 Id: %d Path: %s%s", id, os.TempDir(), fileName)
	// 下载音乐文件
	client := req.C().SetOutputDirectory(os.TempDir())
	_, err = client.R().SetOutputFile(fileName).Get(music.Url)
	if err != nil {
		return
	}
	log.Infof("音乐下载完毕 Id: %d", id)
}
