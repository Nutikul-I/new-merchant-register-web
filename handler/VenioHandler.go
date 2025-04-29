package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"register-service/model"
	"runtime"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type VenioHandler interface {
	GetToken() (model.TokenResponse, error)
	CreateCustomer(payload model.CustomerRequest)
}

type venioHandler struct{}

func NewVenioHandler() VenioHandler {
	return &venioHandler{}
}

var tokenCache = struct {
	sync.RWMutex
	Token string
}{}

func StoreToken(token string) {
	tokenCache.Lock()
	defer tokenCache.Unlock()
	tokenCache.Token = token
}

func GetCachedToken() string {
	tokenCache.RLock()
	defer tokenCache.RUnlock()
	return tokenCache.Token
}

func (vh *venioHandler) GetToken() (model.TokenResponse, error) {
	// Check if a cached token exists and return it
	// if cachedToken := GetCachedToken(); cachedToken != "" {
	//  log.Info("Returning cached token")
	//  return model.TokenResponse{AccessToken: cachedToken}, nil
	// }

	// Define log component
	_, file, _, _ := runtime.Caller(1)
	pc, _, _, _ := runtime.Caller(1)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": file[strings.LastIndex(file, "/")+1:],
		"function":  functionName,
	})

	log.Info("Getting new token from Venio API")

	url := viper.GetString("VENIO_URL") + "/authorization/connect/token"
	payload := strings.NewReader("grant_type=client_credentials&client_id=" +
		viper.GetString("VENIO_CLIENT_ID") + "&client_secret=" + viper.GetString("VENIO_CLIENT_SECRET"))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Errorf("Failed to create request: %v", err)
		return model.TokenResponse{}, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Ocp-Apim-Subscription-Key", viper.GetString("OCP_APIM_SUBSCRIPTION_KEY"))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Errorf("Request failed: %v", err)
		return model.TokenResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Errorf("Failed to read response body: %v", err)
		return model.TokenResponse{}, err
	}

	var venioToken model.TokenResponse
	if err := json.Unmarshal(body, &venioToken); err != nil {
		log.Errorf("Failed to unmarshal token response: %v", err)
		return model.TokenResponse{}, err
	}

	// Store the token for future use
	StoreToken(venioToken.AccessToken)
	log.Info("Stored new token successfully")

	return venioToken, nil
}

func (vh *venioHandler) CreateCustomer(payload model.CustomerRequest) {
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"function":  functionName,
	})

	url := viper.GetString("VENIO_URL") + "/v3/customers"
	method := "POST"

	// Serialize the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Error("Error serializing payload to JSON(Venio): ", err)
		return
	}

	// Create an HTTP request
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(jsonPayload))
	if err != nil {
		log.Error("Error creating HTTP request(Venio): ", err)
		return
	}

	// log.Info("tokenCache.Token: ", tokenCache.Token)

	req.Header.Add("Ocp-Apim-Subscription-Key", viper.GetString("OCP_APIM_SUBSCRIPTION_KEY"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+tokenCache.Token)

	res, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return
	}
	defer res.Body.Close()

	log.Info("Response status from Venio: ", res.Status)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error("Error response from Venio", err)
		return
	}
	log.Info(string(body))
}
