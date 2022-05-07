package routers

import (
	"blive/src/entities"
	"blive/src/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func HandleLiveRouter(pRouter *fiber.Router) {
	router := *pRouter
	// 开始直播
	router.Get("/live/start", func(ctx *fiber.Ctx) error {
		services.CurrentDanmuService.RoomId, _ = strconv.Atoi(ctx.Query("roomId"))
		go services.CurrentDanmuService.Start()
		return ctx.JSON(entities.PureSuccessResult())
	})

	// 结束直播
	router.Get("/live/stop", func(ctx *fiber.Ctx) error {
		services.CurrentDanmuService.Close()
		return ctx.JSON(entities.PureSuccessResult())
	})
}
