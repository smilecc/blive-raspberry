package routers

import (
	"blive/src/database"
	"blive/src/entities"
	"github.com/gobuffalo/packr/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func AppRouter(app *fiber.App) {
	app.Use(filesystem.New(filesystem.Config{
		Root: packr.New("Front Box", "../../front/dist"),
	}))

	apiGroup := app.Group("/api")
	apiGroup.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(entities.SuccessResult("hello world"))
	})

	apiGroup.Get("/config/:name", func(ctx *fiber.Ctx) error {
		var config database.SysConfig
		result := database.DB.Where("name = ?", ctx.Params("name", "")).First(&config)
		if result.Error != nil {
			return ctx.JSON(entities.SuccessResult[*database.SysConfig](nil))
		}

		return ctx.JSON(entities.SuccessResult(config))
	})

	apiGroup.Put("/config/:name", func(ctx *fiber.Ctx) error {
		var configs []database.SysConfig
		database.DB.Where("name = ?", ctx.Params("name", "")).Find(&configs)

		//newConfig := database.SysConfig{}
		return ctx.JSON(entities.SuccessResult[*database.SysConfig](nil))
	})
}
