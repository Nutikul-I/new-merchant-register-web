package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"register-service/model"
	"register-service/service"
	"runtime"
	"sort"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

type NewController interface {
	NewRedirect(c *fiber.Ctx) error
	NewRegister(c *fiber.Ctx) error

	GetNewMerchantRegister(c *fiber.Ctx) error
	CreateNewMerchantRegister(c *fiber.Ctx) error
}

type newController struct {
	newService          service.MRService
	dataAnalyticService service.DataAnalyticService
}

func NewNewController(newService service.MRService, dataAnalyticService service.DataAnalyticService) newController {
	return newController{newService, dataAnalyticService}
}

func (obj *newController) NewRegister(c *fiber.Ctx) error {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	fullURL := c.OriginalURL()
	log.Infof("Full URL: %s", fullURL)

	utm_source := c.Query("utm_source")
	utm_medium := c.Query("utm_medium")
	utm_campaign := c.Query("utm_campaign")
	utm_content := c.Query("utm_content")
	business_category := c.Query("business_category", "")

	//Catagory
	Catagory, err := obj.newService.GetCatagory()
	if err != nil {
		log.Error(err)
	}
	var filtered []model.Category
	for _, cat := range Catagory {
		if cat.NameTH != "--" && cat.NameTH != "ปืน" && cat.NameTH != "อื่นๆ" {
			filtered = append(filtered, cat)
		}
	}

	//SubCatagory
	SubCatagory, err := obj.newService.GetSubCatagory()
	if err != nil {
		log.Error(err)
	}

	lang := c.Query("lang", "TH")

	if lang == "TH" {
		sort.Slice(filtered, func(i, j int) bool {
			return filtered[i].NameTH < filtered[j].NameTH
		})
	} else {
		sort.Slice(filtered, func(i, j int) bool {
			return filtered[i].NameEN < filtered[j].NameEN
		})
	}

	return c.Render("new-register", fiber.Map{
		"Catagory":         filtered,
		"SubCatagory":      SubCatagory,
		"UtmSource":        utm_source,
		"UtmMedium":        utm_medium,
		"UtmCampaign":      utm_campaign,
		"UtmContent":       utm_content,
		"Lang":             lang,
		"BusinessCategory": business_category,
	})
}

func (obj *newController) NewPaySoShopRegister(c *fiber.Ctx) error {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	fullURL := c.OriginalURL()
	log.Infof("Full URL: %s", fullURL)

	utm_source := c.Query("utm_source")
	utm_medium := c.Query("utm_medium")
	utm_campaign := c.Query("utm_campaign")
	utm_content := c.Query("utm_content")
	business_category := c.Query("business_category", "")

	//Catagory
	Catagory, err := obj.newService.GetCatagory()
	if err != nil {
		log.Error(err)
	}
	var filtered []model.Category
	for _, cat := range Catagory {
		if cat.NameTH != "--" && cat.NameTH != "ปืน" && cat.NameTH != "อื่นๆ" {
			filtered = append(filtered, cat)
		}
	}

	//SubCatagory
	SubCatagory, err := obj.newService.GetSubCatagory()
	if err != nil {
		log.Error(err)
	}

	lang := c.Query("lang", "TH")

	if lang == "TH" {
		sort.Slice(filtered, func(i, j int) bool {
			return filtered[i].NameTH < filtered[j].NameTH
		})
	} else {
		sort.Slice(filtered, func(i, j int) bool {
			return filtered[i].NameEN < filtered[j].NameEN
		})
	}

	return c.Render("new-register-paysoshop", fiber.Map{
		"Catagory":         filtered,
		"SubCatagory":      SubCatagory,
		"UtmSource":        utm_source,
		"UtmMedium":        utm_medium,
		"UtmCampaign":      utm_campaign,
		"UtmContent":       utm_content,
		"Lang":             lang,
		"BusinessCategory": business_category,
	})
}

func (obj *newController) NewRedirect(c *fiber.Ctx) error {
	return c.Redirect("/view-register/new/register")
}

func (obj *newController) GetNewMerchantRegister(c *fiber.Ctx) error {
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

	result, err := obj.newService.GetMerchantRegister(input)
	if err != nil {
		log.Errorf("error : %#v", err)
		c.Response().SetStatusCode(http.StatusInternalServerError)
		return c.JSON(response)
	}

	return c.JSON(result)
}

func (obj newController) CreateNewMerchantRegister(c *fiber.Ctx) error {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"function":  functionName,
	})

	enableCaptcha := viper.GetBool("SERVICE_CAPTCHA_ENABLED")

	if enableCaptcha {
		// Parse Turnstile token
		token := c.FormValue("cf-turnstile-response")
		if token == "" {
			log.Error("Turnstile token is missing")
			return c.Status(fiber.StatusInternalServerError).Render("new-register-false", fiber.Map{"Error": "4101"})
		}
		secretKey := viper.GetString("SECREAT_KEY_CAPTCHA")

		tsResponse, err := validateTurnstileToken(secretKey, token)
		if err != nil || !tsResponse.Success {
			log.Errorf("Turnstile validation failed: %v", err)
			return c.Status(fiber.StatusInternalServerError).Render("new-register-false", fiber.Map{"Error": "4101"})
		}
	}

	// Parse request body
	var body model.NewRegisterModel
	if err := c.BodyParser(&body); err != nil {
		log.Errorf("Failed to parse request body: %v", err)
		return c.Status(fiber.StatusInternalServerError).Render("new-register-false", fiber.Map{"Error": "4101"})
	}

	// Proceed with merchant registration
	result, err := obj.newService.CreateNewMerchantRegister(body)
	if err != nil {
		log.Errorf("Error creating merchant register: %v", err)
		return c.Status(fiber.StatusInternalServerError).Render("new-register-false", fiber.Map{"Error": "5101"})
	}

	utm_source := c.Query("utm_source")
	utm_medium := c.Query("utm_medium")
	utm_campaign := c.Query("utm_campaign")
	utm_content := c.Query("utm_content")

	message, err := obj.dataAnalyticService.GetToken()
	if err != nil {
		log.Errorf("Error fetching Venio token: %v", err)
	}
	log.Infof("Token message: %s", message)

	// Create customer in Venio
	obj.dataAnalyticService.CreateCustomer(body, utm_source, utm_medium, utm_campaign, utm_content, "")

	// Uncomment and use this if the service works as expected
	if result == "Register Success" {
		return c.Status(fiber.StatusInternalServerError).Render("new-register-thankyou", fiber.Map{})
	}

	return c.Status(fiber.StatusInternalServerError).Render("new-register-false", fiber.Map{"Error": "4101"})
}

func (obj newController) CreateNewPaySoShopRegister(c *fiber.Ctx) error {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"function":  functionName,
	})

	enableCaptcha := viper.GetBool("SERVICE_CAPTCHA_ENABLED")

	if enableCaptcha {
		// Parse Turnstile token
		token := c.FormValue("cf-turnstile-response")
		if token == "" {
			log.Error("Turnstile token is missing")
			return c.Status(fiber.StatusInternalServerError).Render("new-register-false", fiber.Map{"Error": "4101"})
		}
		secretKey := viper.GetString("SECREAT_KEY_CAPTCHA")

		tsResponse, err := validateTurnstileToken(secretKey, token)
		if err != nil || !tsResponse.Success {
			log.Errorf("Turnstile validation failed: %v", err)
			return c.Status(fiber.StatusInternalServerError).Render("new-register-false", fiber.Map{"Error": "4101"})
		}
	}

	// Parse request body
	var body model.NewRegisterModel
	if err := c.BodyParser(&body); err != nil {
		log.Errorf("Failed to parse request body: %v", err)
		return c.Status(fiber.StatusInternalServerError).Render("new-register-false", fiber.Map{"Error": "4101"})
	}

	// Proceed with merchant registration
	result, err := obj.newService.CreateNewMerchantRegister(body)
	if err != nil {
		log.Errorf("Error creating merchant register: %v", err)
		return c.Status(fiber.StatusInternalServerError).Render("new-register-false", fiber.Map{"Error": "5101"})
	}

	utm_source := c.Query("utm_source")
	utm_medium := c.Query("utm_medium")
	utm_campaign := c.Query("utm_campaign")
	utm_content := c.Query("utm_content")
	assignToOwner := viper.GetString("VENIO_EVENT_PAYSOSHOP_ASSIGN_TO_OWNER")

	message, err := obj.dataAnalyticService.GetToken()
	if err != nil {
		log.Errorf("Error fetching Venio token: %v", err)
	}
	log.Infof("Token message: %s", message)

	// Create customer in Venio
	obj.dataAnalyticService.CreateCustomer(body, utm_source, utm_medium, utm_campaign, utm_content, assignToOwner)

	// Uncomment and use this if the service works as expected
	if result == "Register Success" {
		return c.Status(fiber.StatusInternalServerError).Render("new-register-thankyou", fiber.Map{})
	}

	return c.Status(fiber.StatusInternalServerError).Render("new-register-false", fiber.Map{"Error": "4101"})
}

func validateTurnstileToken(secretKey, token string) (*model.TurnstileResponse, error) {
	// Turnstile verification URL
	const turnstileVerifyURL = "https://challenges.cloudflare.com/turnstile/v0/siteverify"
	// Prepare the request payload
	payload := map[string]string{
		"secret":   secretKey,
		"response": token,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to create payload: %w", err)
	}

	// Make the POST request to Cloudflare's API
	resp, err := http.Post(turnstileVerifyURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to send request to Turnstile: %w", err)
	}
	defer resp.Body.Close()

	// Read and parse the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read Turnstile response: %w", err)
	}

	var tsResponse model.TurnstileResponse
	err = json.Unmarshal(body, &tsResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Turnstile response: %w", err)
	}

	return &tsResponse, nil
}
