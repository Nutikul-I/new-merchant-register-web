package service

import (
	"register-service/handler"
	"register-service/model"
	"runtime"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type DataAnalyticService interface {
	GetToken() (message string, err error)
	CreateCustomer(registermodel model.NewRegisterModel, utm_source, utm_medium, utm_campaign, utm_content, owner string)

	CreateReseller(registermodel model.NewResellerRegisterModel, utm_source, utm_medium, utm_campaign, utm_content, owner string)
	GetTokenByUrlValues() (message string, err error)
}

type dataAnalyticService struct {
	venioHandler handler.VenioHandler
	cachedToken  model.TokenResponse
	tokenExpiry  time.Time
	newService   MRService
}

func NewDataAnalyticService(venioHandler handler.VenioHandler, cachedToken model.TokenResponse, tokenExpiry time.Time, newService MRService) DataAnalyticService {
	return &dataAnalyticService{venioHandler, cachedToken, tokenExpiry, newService}
}

func (obj *dataAnalyticService) GetToken() (string, error) {
	// Define log component
	_, file, _, _ := runtime.Caller(1)
	pc, _, _, _ := runtime.Caller(1)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": file[strings.LastIndex(file, "/")+1:],
		"function":  functionName,
	})

	if obj.cachedToken.AccessToken != "" && time.Now().Before(obj.tokenExpiry.Add(-40*time.Minute)) {
		log.Infof("Get Venio Token from cache, expires at: %v", obj.tokenExpiry.Add(-40*time.Minute))
		return "Using cached token", nil
	}

	log.Info("Generate Venio Token")
	token, err := obj.venioHandler.GetToken()
	if err != nil {
		log.Errorf("Failed to generate token: %v", err)
		obj.cachedToken = model.TokenResponse{}
		return "Failed to gererate token", err
	}

	// Store the new token and set expiry
	obj.cachedToken = token
	obj.tokenExpiry = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)

	return "Got Token", nil
}

func (obj *dataAnalyticService) CreateCustomer(registermodel model.NewRegisterModel, utm_source, utm_medium, utm_campaign, utm_content, owner string) {
	// Define log component
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(1) // Caller(1) gets the function caller
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"function":  functionName,
	})

	log.Infof("Start sending customer data to Venio")

	var interestsName []string
	var categoryNameEN string
	var otherCategory string
	var marketingConsent string
	var owners []string

	if registermodel.CF1148 == "1" {
		interestsName = append(interestsName, "EDC")
	} else {
		interestsName = append(interestsName, "")
	}

	if owner != "" {
		owners = append(owners, owner)
	}

	if registermodel.CF1154 == "1" {
		marketingConsent = "CONSENT_MKT_V1_0"
	} else {
		marketingConsent = ""
	}

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

	for i := 0; i < len(filtered); i++ {
		if registermodel.CategoryNameEN == filtered[i].NameEN {
			categoryNameEN = registermodel.CategoryNameEN
			otherCategory = ""
			break
		} else {
			categoryNameEN = "Others"
			otherCategory = registermodel.Category
		}
	}

	var payload = model.CustomerRequest{
		CustomerName:   registermodel.Name,
		CustomerState:  1,
		CustomerStatus: 0,
		CustomerType:   1,
		SourceName:     "Z-" + categoryNameEN,
		LeadStatus:     1,
		InterestsName:  interestsName,
		CustomFields: []model.CustomField{
			{
				CustomFieldName:  "Website or Social Media",
				CustomFieldValue: []string{registermodel.Website},
			},
			{
				CustomFieldName:  "Accepted Terms & Conditions",
				CustomFieldValue: []string{"CONSENT_T&C_V1_0"},
			},
			{
				CustomFieldName:  "Accepted Consent Marketing",
				CustomFieldValue: []string{marketingConsent},
			},
			{
				CustomFieldName:  "UTM Source",
				CustomFieldValue: []string{utm_source},
			},
			{
				CustomFieldName:  "UTM Medium",
				CustomFieldValue: []string{utm_medium},
			},
			{
				CustomFieldName:  "UTM Campaign",
				CustomFieldValue: []string{utm_campaign},
			},
			{
				CustomFieldName:  "UTM Content",
				CustomFieldValue: []string{utm_content},
			},
			{
				CustomFieldName:  "Business Category - Other",
				CustomFieldValue: []string{otherCategory},
			},
		},
		Contacts: []model.Contact{
			{
				ContactName:   registermodel.Name + " " + registermodel.Surname,
				ContactStatus: true,
				ContactMobile: registermodel.Mobile,
				ContactEmail:  registermodel.Email,
			},
		},
		Owners: owners,
	}

	// payloadJSON, err := json.MarshalIndent(payload, "", "  ")
	// if err != nil {
	// 	log.Fatalf("Error formatting payload: %v", err)
	// }
	// log.Infof("Payload prepared:\n%s", string(payloadJSON))

	log.Info("Send cutomer data to Venio")
	obj.venioHandler.CreateCustomer(payload)

}

func (obj *dataAnalyticService) CreateReseller(registermodel model.NewResellerRegisterModel, utm_source, utm_medium, utm_campaign, utm_content, owner string) {
	// Define log component
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(1) // Caller(1) gets the function caller
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"function":  functionName,
	})

	log.Infof("Start sending customer data to Venio")

	var interestsName []string
	var owners []string

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

	var payload = model.CustomerRequest{
		CustomerName:   registermodel.Name,
		CustomerState:  1,
		CustomerStatus: 0,
		CustomerType:   1,
		SourceName:     "new-reseller-register",
		LeadStatus:     1,
		InterestsName:  interestsName,
		CustomFields: []model.CustomField{
			{
				CustomFieldName:  "Reseller name",
				CustomFieldValue: []string{registermodel.ResellerName},
			},
			{
				CustomFieldName:  "Reseller MID",
				CustomFieldValue: []string{registermodel.ResellerMID},
			},
			{
				CustomFieldName:  "Reseller email",
				CustomFieldValue: []string{registermodel.ResellerEmail},
			},
			{
				CustomFieldName:  "Reseller remark",
				CustomFieldValue: []string{registermodel.ResellerRemark},
			},
		},
		Contacts: []model.Contact{
			{
				ContactName:   registermodel.Name,
				ContactStatus: true,
				ContactPhone:  registermodel.Phone,
				ContactMobile: registermodel.Mobile,
				ContactEmail:  registermodel.Email,
			},
		},
		Owners: owners,
	}

	log.Info("Send cutomer data to Venio")
	obj.venioHandler.CreateCustomer(payload)

}

func (obj *dataAnalyticService) GetTokenByUrlValues() (string, error) {
	// Define log component
	_, file, _, _ := runtime.Caller(1)
	pc, _, _, _ := runtime.Caller(1)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": file[strings.LastIndex(file, "/")+1:],
		"function":  functionName,
	})

	if obj.cachedToken.AccessToken != "" && time.Now().Before(obj.tokenExpiry.Add(-40*time.Minute)) {
		log.Infof("Get Venio Token from cache, expires at: %v", obj.tokenExpiry.Add(-40*time.Minute))
		return "Using cached token", nil
	}

	log.Info("Generate Venio Token")
	token, err := obj.venioHandler.GetTokenByUrlValues()
	if err != nil {
		log.Errorf("Failed to generate token: %v", err)
		obj.cachedToken = model.TokenResponse{}
		return "Failed to gererate token", err
	}

	// Store the new token and set expiry
	obj.cachedToken = token
	obj.tokenExpiry = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)

	log.Info("Generate Venio Token Successfully")
	return "Got Token", nil
}
