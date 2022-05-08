package routers

import (
	"blive/src/entities"
	"blive/src/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func HandleLiveRouter(pRouter *fiber.Router) {
	router := *pRouter

	router.Get("/live/state", func(ctx *fiber.Ctx) error {
		return ctx.JSON(entities.SuccessResult(services.CurrentDanmuService.IsListening()))
	})

	// 开始直播
	router.Get("/live/start", func(ctx *fiber.Ctx) error {
		services.CurrentDanmuService.RoomId, _ = strconv.Atoi(ctx.Query("roomId"))
		go services.CurrentDanmuService.Start()
		go services.CurrentVideoService.StartLiveStream()
		return ctx.JSON(entities.PureSuccessResult())
	})

	// 结束直播
	router.Get("/live/stop", func(ctx *fiber.Ctx) error {
		services.CurrentDanmuService.Close()
		go services.CurrentVideoService.StopLiveStream()
		return ctx.JSON(entities.PureSuccessResult())
	})
}
