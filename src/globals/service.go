package globals

import (
	"os"
	"path"
)

type IService interface {
	Start()
	Close()
}

type IDanmuService interface {
	IService
	IsListening() bool
	GetRoomId() int
	SetRoomId(roomId int)
}

var DanmuService IDanmuService

func GetMusicDir() string {
	dir, _ := os.UserHomeDir()
	dir = path.Join(dir, "blive_tmp/blive_music")
	return dir
}
