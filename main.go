package main

import (
	"blive/src/database"
	"blive/src/routers"
	"blive/src/services"
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
	handleCommand(&danmuChannel, &musicChannel)

	danmuService := services.NewDanmuService(&danmuChannel)
	services.CurrentDanmuService = &danmuService

	musicService := services.NewMusicService(&musicChannel)
	go musicService.Start()

	routers.AppRouter(app)
	log.Fatal(app.Listen(":18000"))

	//file, err := os.OpenFile("D:\\Temp\\test\\a\\a.bat", os.O_CREATE, 0777)
	//if err != nil {
	//	return
	//}
	//defer file.Close()
	//writer := bufio.NewWriter(file)
	//str := "ffmpeg -re -loop 1 -r 3 -t 100 -f image2 -i \"D:\\Temp\\test\\a\\1.jpg\" -i \"D:\\Temp\\test\\a\\1.mp3\" -vf ass=\"D:\\\\Temp\\\\test\\\\a\\\\1.ass\" -pix_fmt yuv420p -crf 24 -preset ultrafast -maxrate 3000k -acodec aac -b:a 192k -c:v h264 -f flv \"D:\\Temp\\test\\a\\output.flv\" -y"
	//writer.WriteString(str)
	//writer.Flush()
	//command := exec.
	//	Command("cmd ")
	//fmt.Println(command.String())
	//command.Stdout = os.Stdout
	//command.Stderr = os.Stderr
	//command.Run()
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
