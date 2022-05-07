package routers

import (
	"blive/src/database"
	"blive/src/entities"
	"github.com/gofiber/fiber/v2"
)

func HandleConfigRouter(pRouter *fiber.Router) {
	router := *pRouter
	// 读取配置项
	router.Get("/config/:name", func(ctx *fiber.Ctx) error {
		var configs []database.SysConfig
		database.DB.Where("name = ?", ctx.Params("name", "")).Find(&configs)
		if len(configs) == 0 {
			return ctx.JSON(entities.SuccessResult[*database.SysConfig](nil))
		}

		return ctx.JSON(entities.SuccessResult(configs[0]))
	})

	// 写入配置项
	router.Put("/config/:name", func(ctx *fiber.Ctx) error {
		configName := ctx.Params("name", "")
		configValue := string(ctx.Body())

		var configs []database.SysConfig
		database.DB.Where("name = ?", configName).Find(&configs)

		var editConfig database.SysConfig

		if len(configs) == 0 {
			editConfig = database.SysConfig{
				Name:  configName,
				Value: configValue,
			}
			database.DB.Create(&editConfig)
		} else {
			editConfig = configs[0]
			editConfig.Value = configValue
			database.DB.Model(&editConfig).Updates(editConfig)
		}

		return ctx.JSON(entities.PureSuccessResult())
	})
}
