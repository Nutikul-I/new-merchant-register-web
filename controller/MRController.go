package controller

import (
	"net/http"
	"register-service/model"
	"register-service/service"
	"runtime"
	"strings"

	"github.com/gofiber/fiber/v2"

	log "github.com/sirupsen/logrus"
)

type MRController interface {
	Ping(c *fiber.Ctx) error

	Redirect(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error

	GetMerchantRegister(c *fiber.Ctx) error
	CreateMerchantRegister(c *fiber.Ctx) error
	ApproveMerchantRegister(c *fiber.Ctx) error
	DenyApproveMerchantRegister(c *fiber.Ctx) error
}

type mrController struct {
	mrService service.MRService
}

func NewMRController(mrService service.MRService) mrController {
	return mrController{mrService}
}

// Ping Healthcheck
func (obj mrController) Ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "register-api 0.99.8"})
}

func (obj mrController) Register(c *fiber.Ctx) error {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	//Catagory
	Catagory, err := obj.mrService.GetCatagory()
	if err != nil {
		log.Error(err)
	}

	//SubCatagory
	SubCatagory, err := obj.mrService.GetSubCatagory()
	if err != nil {
		log.Error(err)
	}

	return c.Render("register", fiber.Map{
		"Catagory":    Catagory,
		"SubCatagory": SubCatagory,
	})
}

func (obj mrController) Redirect(c *fiber.Ctx) error {

	return c.Redirect("/view-register/salesregister")
}

func (obj mrController) GetMerchantRegister(c *fiber.Ctx) error {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	var response []model.TempMerchant

	var input map[string]string
	err := c.BodyParser(&input)

	if err != nil {
		log.Errorf("error : %#v", err)
		c.Response().SetStatusCode(http.StatusInternalServerError)
		return c.JSON(response)
	}

	log.Debugf("input ==> %#v", input)

	result, err := obj.mrService.GetMerchantRegister(input)
	if err != nil {
		log.Errorf("error : %#v", err)
		c.Response().SetStatusCode(http.StatusInternalServerError)
		return c.JSON(response)
	}

	return c.JSON(result)
}

func (obj mrController) CreateMerchantRegister(c *fiber.Ctx) error {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	var body model.RegisterModel
	if err := c.BodyParser(&body); err != nil {
		log.Errorf("error : %#v", err)
		c.Response().SetStatusCode(http.StatusInternalServerError)
		return c.Render("register-false", fiber.Map{})
	}
	log.Infof("body : %#v", body)

	result, err := obj.mrService.CreateMerchantRegister(body)
	if err != nil {
		log.Error(err)
		c.Response().SetStatusCode(http.StatusInternalServerError)
		return c.Render("register-false", fiber.Map{})
	}

	if result == "Register Success" {
		return c.Render("register-thankyou", fiber.Map{})
	}

	return c.Render("register-false", fiber.Map{})
}

func (obj mrController) ApproveMerchantRegister(c *fiber.Ctx) error {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	var response model.ResponseApprove
	response.Status = "Error"

	var input map[string]string
	err := c.BodyParser(&input)

	if err != nil {
		log.Errorf("error : %#v", err)
		c.Response().SetStatusCode(http.StatusInternalServerError)
		response.Message = err.Error()
		return c.JSON(response)
	}

	log.Debugf("input ==> %#v", input)

	result, err := obj.mrService.ApproveMerchantRegister(input, c.IP())
	if err != nil {
		log.Errorf("error : %#v", err)
		c.Response().SetStatusCode(http.StatusInternalServerError)
		response.Message = err.Error()
		return c.JSON(response)
	}

	log.Debugf("result ==> %v", result)

	if result == "Register Success" {
		response.Status = result
		return c.JSON(response)
	}

	c.Response().SetStatusCode(http.StatusInternalServerError)
	return c.JSON(response)
}

func (obj mrController) DenyApproveMerchantRegister(c *fiber.Ctx) error {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	var response model.ResponseApprove
	response.Status = "Error"

	var input map[string]string
	err := c.BodyParser(&input)

	if err != nil {
		log.Errorf("error : %#v", err)
		c.Response().SetStatusCode(http.StatusInternalServerError)
		response.Message = err.Error()
		return c.JSON(response)
	}

	log.Debugf("input ==> %#v", input)

	result, err := obj.mrService.DenyApproveMerchantRegister(input, c.IP())
	if err != nil {
		log.Errorf("error : %#v", err)
		c.Response().SetStatusCode(http.StatusInternalServerError)
		response.Message = err.Error()
		return c.JSON(response)
	}

	log.Debugf("result ==> %v", result)

	if result == "Deny Success" {
		response.Status = result
		return c.JSON(response)
	}

	c.Response().SetStatusCode(http.StatusInternalServerError)
	return c.JSON(response)
}
