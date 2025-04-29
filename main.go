package main

import (
	"register-service/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	log.Info("-= Init Main =-")
}

func main() {
	engine := html.New(viper.GetString("VIEWS_PATH"), ".html")
	log.Debugf("vew path : %v", viper.GetString("VIEWS_PATH"))

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	router.SetupRoutes(app)

	servicePort := viper.GetString("SERVICE_PORT")

	log.Debugf("Service Port: %v", servicePort)
	_ = app.Listen(":" + servicePort)

}
