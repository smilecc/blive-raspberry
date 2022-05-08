package main

import (
	"blive/src/database"
	"blive/src/routers"
	"blive/src/services"
	"blive/src/services/music"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	database.Connect()
	app := fiber.New()
	danmuChannel := make(chan services.DanmuCommand, 10)
	musicChannel := make(chan string, 500)
	encodeChannel := make(chan music.SongDetail, 500)
	handleCommand(&danmuChannel, &musicChannel)

	danmuService := services.NewDanmuService(&danmuChannel)
	services.CurrentDanmuService = &danmuService

	videoService := services.NewVideoService()
	services.CurrentVideoService = &videoService

	musicService := services.NewMusicService(&musicChannel, &encodeChannel)
	go musicService.Start()
	go services.StartEncode(&encodeChannel)

	// 一首构建默认歌曲
	musicChannel <- services.DefaultMusic

	routers.AppRouter(app)
	log.Fatal(app.Listen(":18000"))
}

func handleCommand(danmuChannel *chan services.DanmuCommand, musicChannel *chan string) {
	go func() {
		for {
			command, ok := <-*danmuChannel
			if !ok {
				fmt.Printf("test: false")
				break
			}
			commandJson, _ := json.Marshal(command)
			fmt.Printf("收到指令: %s\n", commandJson)

			if command.CommandName == "点歌" {
				*musicChannel <- command.Arg1
			}
		}
	}()
}
