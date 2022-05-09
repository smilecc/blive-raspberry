package routers

import (
	"blive/src/entities"
	"blive/src/globals"
	"github.com/gobuffalo/packr/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func AppRouter(app *fiber.App) {
	handleWebsocket(app)

	app.Use(cors.New())
	app.Static("/music", globals.GetMusicDir())

	apiGroup := app.Group("/api")

	apiGroup.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(entities.SuccessResult("hello world"))
	})

	HandleConfigRouter(&apiGroup)
	HandleLiveRouter(&apiGroup)

	app.Use(filesystem.New(filesystem.Config{
		Root:       packr.New("Front Box", "../../front/dist"),
		PathPrefix: "/",
	}))
}
