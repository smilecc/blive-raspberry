package routers

import (
	"blive/src/entities"
	"blive/src/globals"
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"net/http"
)

func AppRouter(app *fiber.App, frontFS embed.FS) {
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
		Root:       http.FS(frontFS),
		PathPrefix: "front/dist",
	}))
}
