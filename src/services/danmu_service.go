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

type DanmuService struct {
	RoomId       int
	listening    bool
	blive        *utils.Live
	danmuChannel *chan DanmuCommand
}

type DanmuCommand struct {
	SenderId    int      `json:"senderId"`
	SenderName  string   `json:"senderName"`
	CommandName string   `json:"commandName"`
	SourceDanmu string   `json:"sourceDanmu"`
	Arg1        string   `json:"arg1"`
	Args        []string `json:"args"`
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

func (d *DanmuService) GetRoomId() int {
	return d.RoomId
}

func (d *DanmuService) SetRoomId(roomId int) {
	d.RoomId = roomId
}

func danmuHandler(danmuChannel *chan DanmuCommand) websocket.DanmakuHandler {
	return func(danmu *dto.Danmaku) {
		commands := []string{"点歌"}

		danmuContent := strings.Trim(danmu.Content, " ")
		for _, command := range commands {
			if strings.HasPrefix(danmu.Content, command) {
				arg := strings.Trim(strings.TrimPrefix(danmu.Content, command), " ")
				*danmuChannel <- DanmuCommand{
					SenderId:    danmu.UID,
					SenderName:  danmu.Uname,
					SourceDanmu: danmuContent,
					CommandName: command,
					Arg1:        arg,
				}
			}
		}

		log.Infof("收到弹幕：%s:%s", danmu.Uname, danmu.Content)
	}
}
