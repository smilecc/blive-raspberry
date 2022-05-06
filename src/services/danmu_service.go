package services

import (
	"blive/src/utils"
	"fmt"
	"github.com/botplayerneo/bili-live-api/dto"
	"github.com/botplayerneo/bili-live-api/websocket"
	"strings"
	"sync"
)

var wg sync.WaitGroup

type DanmuService struct {
	blive        *utils.Live
	danmuChannel *chan DanmuCommand
}

type DanmuCommand struct {
	SenderId    int
	SenderName  string
	CommandName string
	SourceDanmu string
	Arg1        string
	Args        []string
}

func NewDanmuService(roomId int, danmuChannel *chan DanmuCommand) DanmuService {
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

func danmuHandler(danmuChannel *chan DanmuCommand) websocket.DanmakuHandler {
	return func(danmu *dto.Danmaku) {
		danmuCommandStrings := strings.Split(danmu.Content, " ")
		if len(danmuCommandStrings) >= 2 {
			*danmuChannel <- DanmuCommand{
				SenderId:    danmu.UID,
				SenderName:  danmu.Uname,
				SourceDanmu: danmu.Content,
				CommandName: danmuCommandStrings[0],
				Arg1:        danmuCommandStrings[1],
				Args:        danmuCommandStrings[1:],
			}
		}

		fmt.Printf("%s:%s\n",
			danmu.Uname,
			danmu.Content,
		)
	}
}
