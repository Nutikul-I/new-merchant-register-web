package router

import (
	"register-service/controller"
	"register-service/handler"
	"register-service/model"
	"register-service/service"
	"runtime"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func SetupRoutes(app *fiber.App) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	dataAnalyticService := service.NewDataAnalyticService(handler.NewVenioHandler(), model.TokenResponse{}, time.Time{}, service.NewMRService(handler.NewMRHandler(), handler.NewVenioHandler(), nil, model.TokenResponse{}, time.Time{}))
	controllermr := controller.NewMRController(service.NewMRService(handler.NewMRHandler(), handler.NewVenioHandler(), dataAnalyticService, model.TokenResponse{}, time.Time{}))
	controllernew := controller.NewNewController(service.NewMRService(handler.NewMRHandler(), handler.NewVenioHandler(), dataAnalyticService, model.TokenResponse{}, time.Time{}), dataAnalyticService)

	api := app.Group("/", func(c *fiber.Ctx) error {
		if !strings.Contains(c.Request().URI().String(), "/ping") {
			log.Infof("all : %v", c.Request().URI().String())
			log.Infof("all : %v", c.OriginalURL())
		}

		return c.Next()
	})
	api.Get("/", controllernew.NewRedirect)

	api.Get("/api/ping", controllermr.Ping)
	api.Get("/api/temp-register", controllermr.GetMerchantRegister)
	api.Post("/api/new-register", controllermr.CreateMerchantRegister)
	api.Put("/api/approve-register", controllermr.ApproveMerchantRegister)
	api.Put("/api/deny-approve-register", controllermr.DenyApproveMerchantRegister)

	api.Get("/salesregister", controllermr.Register)
	api.Post("/new-register", controllermr.CreateMerchantRegister)

	api.Get("/new/register", controllernew.NewRegister)
	api.Post("/new/new-register", controllernew.CreateNewMerchantRegister)

	api.Get("/new/register-paysoshop", controllernew.NewPaySoShopRegister)
	api.Post("/new/new-register-paysoshop", controllernew.CreateNewPaySoShopRegister)

	api.Get("/new/reseller-register", controllernew.NewReseller)
	api.Post("/new/new-reseller-register", controllernew.CreateNewResellerRegister)

}
