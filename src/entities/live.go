package entities

type LiveRoom struct {
	RoomId       int    `json:"roomId"`
	LiveUrl      string `json:"liveUrl"`
	LivePassword string `json:"livePassword"`
}
