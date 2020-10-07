package main

import (
	"cygo_iris/repository"
	"cygo_iris/route"
	"cygo_iris/session"
	"cygo_iris/util/i18n"
	"cygo_iris/util/log"
	"cygo_iris/util/validator"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"os"
)

func main() {
	// Create app.
	app := iris.New()

	// Load setting from ./.env
	_ = godotenv.Load()

	// Set logger.
	// Log Levels => https://github.com/kataras/golog/blob/master/README.md#log-levels
	log.Logger = app.Logger()
	log.Logger.SetLevel(os.Getenv("IRIS_MODE"))

	// Init redis and mongodb.
	session.ConnectRedis()
	if err := repository.ConnectMgo(); err != nil {
		log.Logger.Errorf("Connect to mongodb failed: %s", err)
		panic(err)
	}

	// Init validator.
	app.Validator = validator.NewValidator()

	// Init i18n config.
	app.I18n.DefaultMessageFunc = i18n.DefaultMessageFunc
	if err := app.I18n.Load("./assets/locale/*/*", "en-US", "zh-CN"); err != nil {
		log.Logger.Errorf("Load i18n failed: %s", err)
		panic(err)
	}
	app.I18n.SetDefault("en-US")

	// Init router.
	route.InitRouter(app)

	_ = app.Run(iris.Addr(":" + os.Getenv("PORT")))
}

