package routers

import (
	"blive/src/entities"
	"blive/src/globals"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func HandleLiveRouter(pRouter *fiber.Router) {
	router := *pRouter

	router.Get("/live/state", func(ctx *fiber.Ctx) error {
		return ctx.JSON(entities.SuccessResult(globals.DanmuService.IsListening()))
	})

	// 开始直播
	router.Get("/live/start", func(ctx *fiber.Ctx) error {
		roomId, _ := strconv.Atoi(ctx.Query("roomId"))
		globals.DanmuService.SetRoomId(roomId)

		go globals.DanmuService.Start()
		return ctx.JSON(entities.PureSuccessResult())
	})

	// 结束直播
	router.Get("/live/stop", func(ctx *fiber.Ctx) error {
		globals.DanmuService.Close()
		return ctx.JSON(entities.PureSuccessResult())
	})
}
