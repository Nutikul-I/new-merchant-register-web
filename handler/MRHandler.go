package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"register-service/model"
	"register-service/repository"
	"register-service/util"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

func init() {
	util.Init()
	repository.Init()
}

type MRHandler interface {
	GetCatagory() ([]model.Category, error)
	GetSubCatagory() ([]model.Category, error)

	EncryptTripleDES(password string) (string, error)
	EncryptSecure(password string) (string, error)
	CreateUser(user map[string]string) (string, error)
	AWSSendMail(mail map[string]string) (string, error)
	SendNotiGoogleChat(noti map[string]string) (string, error)
	CreateNewMerchant(data model.NewRegisterModel) (string, error)

	CreateNewReseller(data model.NewResellerRegisterModel) (string, error)
}

type mrHandler struct {
}

func NewMRHandler() mrHandler {
	return mrHandler{}
}

func (obj mrHandler) GetCatagory() ([]model.Category, error) {
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	var data []model.Category

	url := viper.GetString("EXTERNAL_API") + "/customer/getcategory"
	method := "GET"

	log.Debugf("url ==> %s", url)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Errorf("err ==> %v", err)
		return []model.Category{}, err
	}

	req.Header.Add("apikey", viper.GetString("APIKEY"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Errorf("err ==> %v", err)
		return []model.Category{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("err ==> %v", err)
		return []model.Category{}, err
	}

	// fmt.Println(string(body))

	json.Unmarshal(body, &data)

	// log.Debugf("data ==> %#v", data)

	return data, nil
}

func (obj mrHandler) GetSubCatagory() ([]model.Category, error) {
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	var data []model.Category

	url := viper.GetString("EXTERNAL_API") + "/customer/getsubcategory"
	method := "GET"

	log.Debugf("url ==> %s", url)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Errorf("err ==> %v", err)
		return []model.Category{}, err
	}

	req.Header.Add("apikey", viper.GetString("APIKEY"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Errorf("err ==> %v", err)
		return []model.Category{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("err ==> %v", err)
		return []model.Category{}, err
	}

	// fmt.Println(string(body))

	json.Unmarshal(body, &data)

	// log.Debugf("data ==> %#v", data)

	return data, nil
}

func (obj mrHandler) EncryptTripleDES(password string) (string, error) {
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	var encrypt model.EncryptTripleDES
	encrypt_password := "WcdiVfPcYKVw7VX+4mENwA=="

	url := viper.GetString("INTERNAL_API") + "/encrypt/tripledes"
	method := "POST"

	log.Debugf("url ==> %s", url)

	payload := strings.NewReader(`{
		"Message":"` + password + `",
		"KeyPassword":"` + viper.GetString("PASSWORD_ENCRYPT") + `"
	}`)

	log.Debugf("payload ==> %#v", payload)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Errorf("err ==> %v", err)
		return encrypt_password, err
	}

	req.Header.Add("apikey", viper.GetString("APIKEY"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Errorf("err ==> %v", err)
		return encrypt_password, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("err ==> %v", err)
		return encrypt_password, err
	}
	// fmt.Println(string(body))

	json.Unmarshal(body, &encrypt)

	log.Debugf("encrypt ==> %#v", encrypt)

	if encrypt.Status == 200 {
		encrypt_password = encrypt.Data
	}

	return encrypt_password, nil
}

func (obj mrHandler) EncryptSecure(password string) (string, error) {
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	url := viper.GetString("SECURE_API") + "/password?pass=" + password
	method := "POST"

	log.Debugf("url ==> %s", url)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Errorf("err ==> %v", err)
		return "Error", err
	}

	req.Header.Add("apikey", viper.GetString("APIKEY"))

	res, err := client.Do(req)
	if err != nil {
		log.Errorf("err ==> %v", err)
		return "Error", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("err ==> %v", err)
		return "Error", err
	}

	var result model.SecurePassword

	json.Unmarshal(body, &result)

	log.Debugf("result ==> %#v", result)

	return result.Password, nil
}

func (obj mrHandler) CreateUser(user map[string]string) (string, error) {
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	var result model.User

	url := viper.GetString("AUTH_API") + "/users"
	method := "POST"

	log.Debugf("url ==> %s", url)

	payload := strings.NewReader(`{
		"createBy": "` + user["CreateBy"] + `",
		"createDate": "` + user["CreateDate"] + `",
		"entityStatus": true,
		"merchantId": "` + user["MerchantID"] + `",
		"updateBy": "` + user["UpdateBy"] + `",
		"updateDate": "` + user["UpdateDate"] + `",
		"userDesc": "` + user["UserDesc"] + `",
		"userEmail": "` + user["UserEmail"] + `",
		"userName": "` + user["UserName"] + `",
		"userPass": "` + user["UserPassword"] + `"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "0", err
	}
	req.Header.Add("accept", "*/*")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("apikey", "JpywR23@W8")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "0", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "0", err
	}

	json.Unmarshal(body, &result)

	log.Debugf("result ==> %#v", result)

	userId := "0"
	if result.Cid != "" {
		res := strings.Split(result.Links.Self.Href, "/")
		userId = res[4]
	}

	return userId, nil
}

func (obj mrHandler) AWSSendMail(mail map[string]string) (string, error) {
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	url := viper.GetString("AWS_SENDMAIL") + "/api/sendmail?server=AWS"
	method := "POST"

	log.Debugf("url ==> %s", url)

	payload := strings.NewReader(`{
		"MailSubject":"` + mail["Subject"] + `",
		"MailFrom":"` + mail["EmailSupport"] + `",
		"MailTo":"` + mail["ToEmail"] + `",
		"MailCC":"",
		"MailTemplate":"` + mail["TemplateMail"] + `",
		"MailData": [
			{
				"Key": "CustomerName",
				"Value": "` + mail["MerchantCompany"] + `"
			},
			{
				"Key": "CompanyNameEN",
				"Value": "` + mail["MerchantCompany"] + `"
			},
			{
				"Key": "UserName",
				"Value": "` + mail["MerchantUID"] + `"
			},
			{
				"Key": "Password",
				"Value": "` + mail["Password"] + `"
			},
			{
				"Key": "TempID",
				"Value": "` + mail["TempID"] + `"
			},
			{
				"Key": "Name",
				"Value": "` + mail["Name"] + `"
			},
			{
				"Key": "Company",
				"Value": "` + mail["Company"] + `"
			},
			{
				"Key": "Mobile",
				"Value": "` + mail["Mobile"] + `"
			},
			{
				"Key": "Email",
				"Value": "` + mail["Email"] + `"
			},
			{
				"Key": "CategoryTH",
				"Value": "` + mail["CategoryTH"] + `"
			},
			{
				"Key": "CategoryEN",
				"Value": "` + mail["CategoryEN"] + `"
			},
			{
				"Key": "Website",
				"Value": "` + mail["Website"] + `"
			}
		]
	}`)

	log.Debugf("payload ==> %#v", payload)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Errorf("err ==> %v", err)
		return "Error", err
	}

	req.Header.Add("apikey", viper.GetString("APIKEY"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Errorf("err ==> %v", err)
		return "Error", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("err ==> %v", err)
		return "Error", err
	}

	// fmt.Println(string(body))

	var result model.SendMail

	json.Unmarshal(body, &result)

	log.Debugf("result ==> %#v", result)

	return result.Message, nil
}

func (obj mrHandler) SendNotiGoogleChat(noti map[string]string) (string, error) {
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	url := viper.GetString("GOOGLECHANT_API")
	method := "POST"

	log.Debugf("url ==> %s", url)

	payload := strings.NewReader(`{"text": "[Alert] -== New Merchant Registe ==- \nTempID : ` + noti["TempID"] + ` , \nName : ` + noti["Name"] + ` , \nCompany : ` + noti["Company"] + ` , \nEmail : ` + noti["Email"] + ` , \nMobile : ` + noti["Mobile"] + ` , \nWebsite : ` + noti["Website"] + `"}`)

	log.Debugf("payload ==> %#v", payload)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Errorf("err ==> %v", err)
		return "Error", err
	}

	req.Header.Add("Content-Type", "text/plain")

	res, err := client.Do(req)
	if err != nil {
		log.Errorf("err ==> %v", err)
		return "Error", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("err ==> %v", err)
		return "Error", err
	}

	fmt.Println(string(body))

	return "COMPLETE", nil
}

func (obj mrHandler) CreateNewMerchant(data model.NewRegisterModel) (string, error) {
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	urlStr := viper.GetString("NEW_REGISTER_VTIGER")
	method := "POST"

	currentDate := time.Now().Format("2006-01-02")

	formData := url.Values{}
	formData.Set("__vtrftk", viper.GetString("VTRFTK"))
	formData.Set("publicid", viper.GetString("PUBLICID"))
	formData.Set("urlencodeenable", "1")
	formData.Set("name", "LDP-FB Group")
	formData.Set("mobile", data.Mobile)
	formData.Set("email", data.Email)
	formData.Set("industry", data.Category)
	formData.Set("leadsource", data.LeadSource)
	formData.Set("leadstatus", "New Lead")
	formData.Set("company", data.Company)
	formData.Set("website", data.Website)
	formData.Set("cf_1129", data.CF1129)
	formData.Set("cf_1148", data.CF1148)
	formData.Set("cf_1150", data.CF1150)
	formData.Set("cf_1152", currentDate)
	formData.Set("cf_1154", data.CF1154)
	formData.Set("cf_1156", currentDate)
	parts := strings.Fields(data.Name)
	if len(parts) > 1 {
		formData.Set("firstname", parts[0])
		formData.Set("lastname", strings.Join(parts[1:], " "))
	} else {
		formData.Set("firstname", data.Name)
		formData.Set("lastname", "-")
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, urlStr, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Setting headers for form submission
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
		return "ERROR", nil
	}
	defer resp.Body.Close()

	log.Printf("Response status: %s", resp.Status)

	return "COMPLETE", nil
}

func (obj mrHandler) CreateNewReseller(data model.NewResellerRegisterModel) (string, error) {
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	urlStr := viper.GetString("NEW_REGISTER_VTIGER")
	method := "POST"

	// currentDate := time.Now().Format("2006-01-02")

	formData := url.Values{}
	formData.Set("__vtrftk", viper.GetString("VTRFTK"))
	formData.Set("publicid", viper.GetString("PUBLICID"))
	formData.Set("urlencodeenable", "1")
	formData.Set("name", "LDP-FB Group")
	formData.Set("mobile", data.Mobile)
	formData.Set("email", data.Email)
	formData.Set("category", data.Category)
	formData.Set("categorynameTH", data.CategoryNameTH)
	formData.Set("categorynameEN", data.CategoryNameEN)
	formData.Set("leadsource", data.LeadSource)
	formData.Set("leadstatus", "New Lead")
	formData.Set("cf1148", data.ServiceEDC)
	formData.Set("reseller-MID", data.ResellerMID)
	formData.Set("reseller-email", data.ResellerEmail)
	formData.Set("reseller-remark", data.ResellerRemark)
	formData.Set("reseller-name", data.ResellerName)
	formData.Set("cf1150", data.AcceptPrivacyPolicy)
	formData.Set("cf1154", data.AcceptMarketingConsent)
	formData.Set("leadsource", data.LeadSource)
	formData.Set("company", data.Company)
	formData.Set("website", data.Website)
	formData.Set("cf_1129", data.CF1129)
	parts := strings.Fields(data.Name)
	if len(parts) > 1 {
		formData.Set("firstname", parts[0])
		formData.Set("lastname", strings.Join(parts[1:], " "))
	} else {
		formData.Set("firstname", data.Name)
		formData.Set("lastname", "-")
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, urlStr, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Setting headers for form submission
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
		return "ERROR", nil
	}
	defer resp.Body.Close()

	log.Printf("Response status: %s", resp.Status)

	return "COMPLETE", nil
}
