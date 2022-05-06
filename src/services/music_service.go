package services

import (
	musicFactory "blive/src/services/music"
	"encoding/json"
	"fmt"
	"github.com/botplayerneo/bili-live-api/log"
	"regexp"
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

			log.Infof("收到点歌指令: %s", music)
			isNumber, _ := regexp.MatchString("\\d+", music)
			musicService := musicFactory.GetMusicService(musicFactory.NeteaseService)
			var songDetail *musicFactory.SongDetail
			if isNumber {
				songDetail, _ = musicService.GetSongById(music)
			} else {
				songDetail, _ = musicService.GetSongByName(music)
			}

			if songDetail != nil {
				songJson, _ := json.Marshal(songDetail)
				log.Infof("获取到歌曲详情：%s", songJson)
			}
		}

	}()
	wg.Wait()
}
