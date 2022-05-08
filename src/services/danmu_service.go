package services

import (
	"blive/src/utils"
	"github.com/botplayerneo/bili-live-api/dto"
	"github.com/botplayerneo/bili-live-api/log"
	"github.com/botplayerneo/bili-live-api/websocket"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var CurrentDanmuService *DanmuService

type DanmuService struct {
	RoomId       int
	listening    bool
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

func NewDanmuService(danmuChannel *chan DanmuCommand) DanmuService {
	service := DanmuService{0, false, nil, danmuChannel}
	return service
}

func (d *DanmuService) Start() {
	if d.listening {
		return
	}

	log.Info("开始监听直播间弹幕")
	d.listening = true
	d.blive = utils.NewBLive(d.RoomId)
	wg.Add(1)
	go func() {
		d.blive.RegisterHandlers(
			danmuHandler(d.danmuChannel),
		)
		d.blive.Start()
	}()
	wg.Wait()
	d.listening = false
}

func (d *DanmuService) IsListening() bool {
	return d.listening
}

func (d *DanmuService) Close() {
	if d.listening {
		d.blive.Close()
		wg.Done()
		log.Info("停止监听直播间弹幕")
	}
}

func danmuHandler(danmuChannel *chan DanmuCommand) websocket.DanmakuHandler {
	return func(danmu *dto.Danmaku) {
		danmuCommandStrings := strings.Split(strings.Trim(danmu.Content, " "), " ")
		if len(danmuCommandStrings) >= 2 {
			*danmuChannel <- DanmuCommand{
				SenderId:    danmu.UID,
				SenderName:  danmu.Uname,
				SourceDanmu: danmu.Content,
				CommandName: strings.Trim(danmuCommandStrings[0], " "),
				Arg1:        strings.Trim(danmuCommandStrings[1], " "),
				Args:        danmuCommandStrings[1:],
			}
		}

		log.Infof("收到弹幕：%s:%s", danmu.Uname, danmu.Content)
	}
}
