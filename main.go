package main

import (
	"blive/src/database"
	"blive/src/globals"
	"blive/src/routers"
	"blive/src/services"
	"blive/src/services/music"
	"embed"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

//go:embed front/dist/*
var frontFS embed.FS

func main() {
	database.Connect()
	app := fiber.New()
	danmuChannel := make(chan services.DanmuCommand, 10)
	musicChannel := make(chan string, 500)
	encodeChannel := make(chan music.SongDetail, 500)
	handleCommand(&danmuChannel, &musicChannel)

	danmuService := services.NewDanmuService(&danmuChannel)
	globals.DanmuService = &danmuService

	musicService := services.NewMusicService(&musicChannel, &encodeChannel)
	go musicService.Start()

	routers.AppRouter(app, frontFS)
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

			routers.SendWebsocketBroadcast(routers.WebsocketMessage[services.DanmuCommand]{"danmu_command", command})
		}
	}()
}
