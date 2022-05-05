package services

import (
	"blive/src/utils"
	"fmt"
	"github.com/botplayerneo/bili-live-api/dto"
	"github.com/botplayerneo/bili-live-api/websocket"
	"sync"
)

var wg sync.WaitGroup

type DanmuService struct {
	blive *utils.Live
	danmuChannel *chan string
}

func NewDanmuService(roomId int, danmuChannel *chan string) DanmuService {
	service := DanmuService{utils.NewBLive(roomId), danmuChannel}
	return service
}

func (d *DanmuService) Start() {
	wg.Add(1)
	go func() {
		d.blive.RegisterHandlers(
			danmuHandler(d.danmuChannel),
		)
		d.blive.Start()
	}()
	wg.Wait()
}

func (d *DanmuService) Close() {
	d.blive.Close()
	wg.Done()
}

func danmuHandler(danmuChannel *chan string) websocket.DanmakuHandler {
	return func(danmu *dto.Danmaku) {
		*danmuChannel <- danmu.Content
		fmt.Printf("%s:%s\n",
			danmu.Uname,
			danmu.Content,
		)
	}
}
