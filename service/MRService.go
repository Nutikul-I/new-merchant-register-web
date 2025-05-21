package service

import (
	"errors"
	"math"
	"register-service/handler"
	"register-service/model"
	"register-service/repository"
	"register-service/util"
	"runtime"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type MRService interface {
	GetCatagory() ([]model.Category, error)
	GetSubCatagory() ([]model.Category, error)

	GetMerchantRegister(data map[string]string) (model.TempMerchant, error)
	CreateMerchantRegister(registermodel model.RegisterModel) (string, error)
	CreateNewMerchantRegister(registermodel model.NewRegisterModel) (string, error)
	ApproveMerchantRegister(data map[string]string, ipAddress string) (string, error)
	DenyApproveMerchantRegister(data map[string]string, ipAddress string) (string, error)

	CreateNewResellerRegister(registermodel model.NewResellerRegisterModel) (string, error)
}

type mrService struct {
	mrHandler           handler.MRHandler
	venioHandler        handler.VenioHandler
	dataAnalyticService DataAnalyticService
	cachedToken         model.TokenResponse // Store the cached token
	tokenExpiry         time.Time
}

func NewMRService(mrHandler handler.MRHandler, venioHandler handler.VenioHandler, dataAnalyticService DataAnalyticService, cachedToken model.TokenResponse, tokenExpiry time.Time) *mrService {
	return &mrService{mrHandler, venioHandler, dataAnalyticService, cachedToken, tokenExpiry}
}

func (obj *mrService) GetCatagory() ([]model.Category, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	log.Infof("-= Get Catatory =-")
	category, err := obj.mrHandler.GetCatagory()
	if err != nil {
		log.Errorf("Error category ==> %#v", err)
	}

	return category, nil
}

func (obj *mrService) GetSubCatagory() ([]model.Category, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	log.Infof("-= Get SubCatatory =-")
	subcategory, err := obj.mrHandler.GetSubCatagory()
	if err != nil {
		log.Errorf("Error category ==> %#v", err)
	}

	return subcategory, nil
}

func (obj *mrService) GetMerchantRegister(data map[string]string) (model.TempMerchant, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	log.Debugf("data ==> %#v", data)

	var response model.TempMerchant

	if data["enddate"] == "" && data["startdate"] == "" {
		return response, nil
	}

	// get counttemp
	count, err := repository.CountTempMerchantRegister(data)
	if err != nil {
		log.Errorf("DB Error : %#v", err)
		return response, err
	}

	log.Debugf("count ==> %#v", count)

	if count == 0 {
		return response, nil
	}

	// get temp
	temp, err := repository.GetTempMerchantRegister(data)
	if err != nil {
		log.Errorf("DB Error : %#v", err)
		return response, err
	}

	// log.Debugf("temp ==> %#v", temp)

	var limit = 20
	var total_page = 0.0
	var next_page = 0
	page, _ := strconv.Atoi(data["page"])
	if data["limit"] != "" {
		limitint, _ := strconv.Atoi(data["limit"])
		limit = limitint
	}
	total_page = (float64(count) / float64(limit))
	x_page := math.Ceil(total_page)
	if total_page == 0 {
		total_page = 1
	}
	next_page = page + 1
	if page == int(x_page) {
		next_page = 0
	}

	temp.Page = page
	temp.TotalData = count
	temp.TotalPage = int(x_page)
	temp.NextPage = next_page

	return temp, nil
}

func (obj *mrService) CreateMerchantRegister(registermodel model.RegisterModel) (string, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	// log.Debugf("input : %v %v %v %v ", registermodel.Name, registermodel.Company, registermodel.Email, registermodel.Mobile)

	// check input
	if len(registermodel.Email) <= 0 || len(registermodel.Company) <= 0 || len(registermodel.Name) <= 0 || len(registermodel.Website) <= 0 {
		log.Errorf("Error : Incorrect information")
		return "Error", errors.New("Incorrect information")
	}

	// check merchantconpayny
	company_status, err := repository.CheckMerchantCompany(registermodel.Company)
	if err != nil {
		log.Errorf("DB Error : %#v", err)
		return "Error", err
	}

	if !company_status {
		log.Errorf("DB Error : Duplicate Merchant Company")
		return "Error", errors.New("Duplicate Merchant Company")
	}

	// datetime
	startDate := time.Now().Local()
	startTime := startDate.Format("2006-01-02 15:04:05")

	// log.Debug("-= Can be created. No information found. =-")

	log.Infof("-= Set Data =-")
	merchant := make(map[string]string)
	merchant["ServiceTypeCode"] = "L"
	merchant["Name"] = registermodel.Name
	merchant["Surname"] = registermodel.Surname
	merchant["Company"] = registermodel.Company
	merchant["Address"] = registermodel.Address
	merchant["Address2"] = registermodel.Address2
	merchant["City"] = registermodel.City
	merchant["State"] = registermodel.State
	merchant["Postalcode"] = registermodel.Postalcode
	merchant["CountryCode"] = "TH"
	merchant["Tel"] = registermodel.Tel
	merchant["Mobile"] = registermodel.Mobile
	merchant["Fax"] = registermodel.Fax
	merchant["Email"] = registermodel.Email
	merchant["Website"] = registermodel.Website
	merchant["SoSimple"] = "0"
	merchant["LangCode"] = registermodel.Lang
	merchant["CreateDateTime"] = startTime
	merchant["CategoryNameTH"] = registermodel.CategoryNameTH
	merchant["CategoryNameEN"] = registermodel.CategoryNameEN
	merchant["LeadNo"] = registermodel.LeadNo

	merchant["Agency"] = registermodel.Agency
	if registermodel.Agency == "" {
		merchant["Agency"] = "-"
	}
	if registermodel.Lang == "" {
		merchant["LangCode"] = "TH"
	}
	merchant["ProvinceID"] = strconv.Itoa(registermodel.ProvinceID)
	if strconv.Itoa(registermodel.ProvinceID) == "0" {
		merchant["ProvinceID"] = "73"
	}
	if registermodel.Tel == "" {
		merchant["Tel"] = "-"
	}
	if registermodel.Mobile == "" {
		merchant["Mobile"] = "-"
	}
	merchant["CatID"] = registermodel.Category
	if registermodel.Category == "" {
		merchant["CatID"] = "0"
	}
	merchant["SCatID"] = registermodel.Subcategory
	if registermodel.Subcategory == "" {
		merchant["SCatID"] = "0"
	}

	log.Debugf("merchant ==> %#v", merchant)

	// insert temp merchant
	tempmerchant, err := repository.CreateTempMerchant(merchant)
	if err != nil {
		log.Errorf("DB Error : %#v", err)
		return "Error", err
	}

	log.Debugf("create temp merchants ==> %#v", tempmerchant)

	if tempmerchant != "COMPLETE" {
		log.Errorf("DB Error : Can't not create temp merchants")
		return "Error", errors.New("Can't not create temp merchants")
	}

	// get tempid
	tempid, err := repository.GetTempMerchantId(merchant["Company"])
	if err != nil {
		log.Errorf("DB Error : %#v", err)
		return "Error", err
	}

	log.Debugf("tempid ==> %#v", tempid)

	if tempid == 0 {
		log.Errorf("Error : No tempid information found.")
		return "Error", errors.New("No tempid information found.")
	}

	merchant["TempID"] = strconv.Itoa(tempid)
	merchant["BankCode"] = registermodel.BankCode
	if registermodel.BankCode == "" {
		merchant["BankCode"] = "SCB"
	}
	merchant["AccountNo"] = registermodel.AccountNo
	merchant["AccountName"] = registermodel.AccountName
	merchant["Branch"] = registermodel.Branch
	merchant["AccountTypeCode"] = registermodel.Accounttype
	if registermodel.Accounttype == "" {
		merchant["AccountTypeCode"] = "S"
	}
	merchant["UpCountryStatus"] = registermodel.Upcountry
	if registermodel.Upcountry == "" {
		merchant["UpCountryStatus"] = "0"
	}

	// insert temp merchant bank account
	tempbankmerchant, err := repository.CreateTempBankMerchant(merchant)
	if err != nil {
		log.Errorf("DB Error : %#v", err)
		return "Error", err
	}

	log.Debugf("create temp bank merchants ==> %#v", tempbankmerchant)

	if tempbankmerchant != "COMPLETE" {
		log.Errorf("DB Error : Can't not create temp bank merchants")
		return "Error", errors.New("Can't not create temp bank merchants")
	}

	// Need new template email alert create tempmerchant cp
	log.Infof("-= Send Mail =-")
	mail := make(map[string]string)
	mail["ToEmail"] = merchant["Email"]
	mail["TempID"] = merchant["TempID"]
	mail["Name"] = merchant["Name"]
	mail["Company"] = merchant["Company"]
	mail["Mobile"] = merchant["Mobile"]
	mail["Email"] = merchant["Email"]
	mail["CategoryTH"] = merchant["CategoryNameTH"]
	mail["CategoryEN"] = merchant["CategoryNameEN"]
	mail["Website"] = merchant["Website"]
	mail["EmailSupport"] = viper.GetString("EMAIL_SUPPORT")
	env := viper.GetString("SERVICE_ENV")
	log.Debugf("env ==> %#v", env)
	if env != "Production" {
		mail["Subject"] = "[Non-Prod] ยืนยันการสมัครและเตรียมเอกสารเพื่อเปิดใช้บริการ Payment Solutions | Confirmation of Registration and Preparation of Documents for PaySolutions Service."
	} else {
		mail["Subject"] = "ยืนยันการสมัครและเตรียมเอกสารเพื่อเปิดใช้บริการ Payment Solutions | Confirmation of Registration and Preparation of Documents for PaySolutions Service."
	}
	mail["TemplateMail"] = viper.GetString("TEMPLATE_EMAIL_WAIT_APPROVE")

	sendmail, err := obj.mrHandler.AWSSendMail(mail)
	if err != nil {
		log.Errorf("Error sendmail ==> %#v", err)
	}

	log.Debugf("Send Mail ==> %#v", sendmail)

	log.Infof("-= Send noti to google chat =-")
	noti := make(map[string]string)
	noti["TempID"] = merchant["TempID"]
	noti["Name"] = merchant["Name"]
	noti["Company"] = merchant["Company"]
	noti["Email"] = merchant["Email"]
	noti["Mobile"] = merchant["Mobile"]
	noti["Website"] = merchant["Website"]

	sendnoti, err := obj.mrHandler.SendNotiGoogleChat(noti)
	if err != nil {
		log.Errorf("Error sendnoti ==> %#v", err)
	}

	log.Debugf("Send Noti Google Chat ==> %#v", sendnoti)

	log.Infof("-= Merchant ==> %v , CreateDate ==> %v =-", merchant["Company"], startTime)

	return "Register Success", nil
}

func (obj mrService) CreateNewMerchantRegister(registermodel model.NewRegisterModel) (string, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	// check input
	if len(registermodel.Email) <= 0 || len(registermodel.Company) <= 0 || len(registermodel.Name) <= 0 || len(registermodel.Mobile) <= 0 {
		log.Errorf("Error : Incorrect information")
		return "Error", errors.New("Incorrect information")
	}

	// insert temp merchant
	tempmerchant, err := obj.mrHandler.CreateNewMerchant(registermodel)
	if err != nil {
		log.Errorf("DB Error : %#v", err)
		return "Error", err
	}

	log.Debugf("create temp merchants ==> %#v", tempmerchant)

	if tempmerchant != "COMPLETE" {
		return "Error", errors.New("Missing invalid field")
	}
	return "Register Success", nil
}

func (obj *mrService) ApproveMerchantRegister(data map[string]string, ipAddress string) (string, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	log.Debugf("data ==> %#v ", data)

	// check input
	if len(data["tempid"]) <= 0 || len(data["company"]) <= 0 {
		log.Errorf("Error : Incorrect information")
		return "Error", errors.New("Incorrect information")
	}

	// get TempsMerchants
	tempmerchant, err := repository.GetTempMerchant(data["tempid"], data["company"])
	if err != nil {
		log.Errorf("DB Error : %#v", err)
		return "Error", err
	}

	if data["tempid"] != tempmerchant.TempID {
		log.Errorf("Error : Incorrect information")
		return "Error", errors.New("Incorrect information")
	}

	log.Debugf("tempmerchant ==> %#v", tempmerchant)

	log.Infof("-= Set Data =-")
	merchant := make(map[string]string)
	merchant["ServiceTypeCode"] = "L"
	merchant["Name"] = tempmerchant.Name
	merchant["Surname"] = tempmerchant.Surname
	merchant["Company"] = tempmerchant.Company
	merchant["Address"] = "-"
	merchant["Address2"] = "-"
	merchant["City"] = "-"
	merchant["State"] = "-"
	merchant["Postalcode"] = "-"
	merchant["CountryCode"] = "TH"
	merchant["Tel"] = "-"
	merchant["Mobile"] = tempmerchant.Mobile
	merchant["Fax"] = "-"
	merchant["Email"] = tempmerchant.Email
	merchant["Website"] = tempmerchant.Website
	merchant["SoSimple"] = "0"
	merchant["ProvinceID"] = "73"
	merchant["Tel"] = "-"
	merchant["LeadNo"] = tempmerchant.LeadNo

	merchant["LangCode"] = tempmerchant.LangCode
	if tempmerchant.LangCode == "" {
		merchant["LangCode"] = "TH"
	}

	merchant["Mobile"] = tempmerchant.Mobile
	if tempmerchant.Mobile == "" {
		merchant["Mobile"] = "-"
	}

	merchant["CatID"] = tempmerchant.CatID
	if tempmerchant.CatID == "" {
		merchant["CatID"] = "0"
	}
	merchant["SCatID"] = tempmerchant.SCatID
	if tempmerchant.SCatID == "" {
		merchant["SCatID"] = "0"
	}

	merchant["TempID"] = tempmerchant.TempID
	merchant["BankCode"] = "SCB"
	merchant["AccountNo"] = ""
	merchant["AccountName"] = ""
	merchant["Branch"] = ""
	merchant["AccountTypeCode"] = ""
	merchant["AccountTypeCode"] = "S"
	merchant["UpCountryStatus"] = "0"

	log.Debugf("merchant ==> %#v", merchant)

	// // insert temp merchant bank account
	tempbankmerchant, err := repository.CreateTempBankMerchant(merchant)
	if err != nil {
		log.Errorf("DB Error : %#v", err)
		return "Error", err
	}

	log.Debugf("create temp bank merchants ==> %#v", tempbankmerchant)

	if tempbankmerchant != "COMPLETE" {
		log.Errorf("DB Error : Can't not create temp bank merchants")
		return "Error", errors.New("Can't not create temp bank merchants")
	}

	// datetime
	startDate := time.Now().Local()
	startTime := startDate.Format("2006-01-02 15:04:05")
	endDate := time.Now().AddDate(5, 0, 0)
	endTime := endDate.Format("2006-01-02 15:04:05")

	// Create Merchant don't approve
	log.Infof("-= Create Merchant don't Approve =-")
	// gen password
	log.Infof("-= gen password =-")
	password := util.GenPassword(8)
	log.Debugf("password ==> %#v", password)
	log.Infof("-= encrypt tripledes =-")
	encrypt_password, err := obj.mrHandler.EncryptTripleDES(password)
	if err != nil {
		log.Errorf("Encrypt password : %#v", err)
		merchant["Password"] = "WcdiVfPcYKVw7VX+4mENwA=="
	}

	log.Debugf("encrypt_password ==> %#v", encrypt_password)
	merchant["Password"] = encrypt_password
	merchant["StartDate"] = startTime
	merchant["ExpiredDate"] = endTime
	merchant["CurrencyCode"] = "00"
	merchant["ServiceTypeCode"] = "L"
	merchant["MonthlyFee"] = "0"
	merchant["VisaAccepted"] = "1"
	merchant["MasterCardAccepted"] = "1"
	merchant["JCBAccepted"] = "1"
	merchant["AMEXAccepted"] = "1"
	merchant["ChargePercent"] = "3.75"
	merchant["AMEXChargePercent"] = "4.75"
	merchant["BillingAddressEnable"] = "0"
	merchant["DemoStatus"] = "1"
	merchant["AutoRenew"] = "0"
	merchant["AccountEnable"] = "1"
	merchant["TransferMoneyPeriodCode"] = "S"

	// insert Merchant
	log.Infof("-= Create Merchant =-")
	Merchant, err := repository.CreateMerchant(merchant)
	if err != nil {
		log.Errorf("DB Error : %#v", err)
		return "Error", err
	}

	log.Debugf("Create Merchant ==> %#v", Merchant)

	if Merchant != "COMPLETE" {
		log.Errorf("DB Create Merchant Error : %#v", err)
		return "Error", err
	}

	// get MerchantID
	merchantid, err := repository.GetMerchantID(merchant)

	if err != nil {
		log.Errorf("DB Get MerchantID Error : %#v", err)
		return "Error", err
	}

	uid := util.GenMerchantUID(3)

	merchantuid := uid + merchantid

	merchant["MerchantID"] = merchantid
	merchant["MerchantUID"] = merchantuid

	log.Debugf("merchant['MerchantID'] ==> %#v", merchant["MerchantID"])
	log.Debugf("merchant['MerchantUID'] ==> %#v", merchant["MerchantUID"])

	//update MerchantUID
	log.Infof("-= Update MerchantUID =-")
	MerchantUID, err := repository.UpdateMerchantUID(merchant)
	if err != nil {
		log.Errorf("DB Error : %#v", err)
		return "Error", err
	}

	log.Debugf("Update MerchantUID ==> %#v", MerchantUID)

	// create merchant configuration
	log.Infof("-= Create Merchant Configuration =-")
	merchantconfig, err := repository.CreateMerchantConfiguration(merchant)
	if err != nil {
		log.Errorf("DB Error : %#v", err)
		return "Error", err
	}

	log.Debugf("Create Merchant Configuration ==> %#v", merchantconfig)

	// create merchant counterservice
	merchant["Enabled"] = "1"
	merchant["ServiceTypeCodeCS"] = "LCS"

	counterservice, err := repository.CreateMerchantCounterService(merchant)
	if err != nil {
		log.Errorf("DB Error : %#v", err)
		return "Error", err
	}

	log.Debugf("Create CounterService ==> %#v", counterservice)

	// add agency
	merchant["Agency"] = tempmerchant.Agency
	if tempmerchant.Agency != "" {
		Reseller, err := repository.CreateMerchantReseller(merchant)
		if err != nil {
			log.Errorf("DB Error ==> %#v", err)
			return "Error", err
		}

		log.Debugf("Create Merchant Reseller ==> %#v", Reseller)
	}

	log.Infof("-= Create Channel =-")
	// get channel
	channel, _ := util.GetChannel()

	log.Debugf("Channel Merchant ==> %#v", channel)

	for _, s := range channel {

		channelmerchant, err := repository.CreateChannelMerchant(merchant, s)
		if err != nil {
			log.Errorf("DB Error : %#v", err)
			return "Error", err
		}

		log.Debugf("Create ChannelMerchant ==> %#v : %#v", s.ChannelCode, channelmerchant)
	}

	log.Infof("-= Create User =-")

	// encrypt secure java
	merchant["PasswordSecure"] = "ez6IEvdCfOyQThb1mB5WF_n_J6kKdso6KEw_k_HMbQ_j__j_"
	encrypt_secure, err := obj.mrHandler.EncryptSecure(password)
	if err != nil {
		log.Errorf("Error encrypt_secure ==> %#v", err)
	}
	if encrypt_secure != "Error" {
		merchant["PasswordSecure"] = encrypt_secure
	}

	log.Debugf("PasswordSecure ==> %#v", encrypt_secure)

	// add user
	user := make(map[string]string)
	user["MerchantID"] = merchant["MerchantID"]
	user["UserDesc"] = merchant["Company"]
	user["UserEmail"] = merchant["Email"]
	user["UserName"] = merchant["MerchantUID"]
	user["UserPassword"] = merchant["PasswordSecure"]
	user["CreateDate"] = startTime
	user["UpdateDate"] = startTime
	user["CreateBy"] = "System"
	user["UpdateBy"] = "System"

	userID := "0"
	create_user, err := obj.mrHandler.CreateUser(user)
	if err != nil {
		log.Errorf("Error create user ==> %#v", err)
		user, err := repository.GetUser(merchant)
		if err != nil {
			log.Errorf("DB Error : %#v", err)
		}
		userID = user
	}
	if create_user != "0" {
		userID = create_user
	}

	user["UserID"] = userID
	user["UserRole"] = "21"
	log.Debugf("Create UserID ==> %#v", user["UserID"])

	// add role
	user_role, err := repository.CreateRole(user)
	if err != nil {
		log.Errorf("DB Error : %#v", err)
	}

	log.Debugf("Create User Role ==> %#v", user_role)

	message := "-"
	if len(data["message"]) > 0 {
		message = data["message"]
	}

	username := "-"
	if len(data["username"]) > 0 {
		message = data["username"]
	}

	// update tempsmerchants
	uptemp, err := repository.UpdateTempMerchants(data["tempid"], data["company"], message, merchant["MerchantUID"], username)
	if err != nil {
		log.Errorf("DB Error : %#v", err)
	}

	log.Debugf("Update TempMerchant ==> %#v", uptemp)

	log.Infof("-= Send Mail =-")
	mail := make(map[string]string)
	mail["ToEmail"] = merchant["Email"]
	mail["MerchantUID"] = merchant["MerchantUID"]
	mail["Password"] = password
	mail["MerchantCompany"] = merchant["Company"]
	mail["EmailSupport"] = viper.GetString("EMAIL_SUPPORT")
	mail["TemplateMail"] = viper.GetString("TEMPLATE_EMAIL_APPROVE")

	env := viper.GetString("SERVICE_ENV")
	log.Debugf("env ==> %#v", env)
	if env != "Production" {
		mail["Subject"] = "[Non-Prod] การสมัครระบบชำระเงินออนไลน์ Pay Solutions (สถานะ-ยืนยันการสมัคร) พร้อมใช้งาน Merchant ID : " + merchant["MerchantUID"] + " | Pay Solutions online payment system application (status - confirmation of application) ready to use Merchant ID : " + merchant["MerchantUID"]
	} else {
		mail["Subject"] = "การสมัครระบบชำระเงินออนไลน์ Pay Solutions (สถานะ-ยืนยันการสมัคร) พร้อมใช้งาน Merchant ID : " + merchant["MerchantUID"] + " | Pay Solutions online payment system application (status - confirmation of application) ready to use Merchant ID : " + merchant["MerchantUID"]
	}

	sendmail, err := obj.mrHandler.AWSSendMail(mail)
	if err != nil {
		log.Errorf("Error sendmail ==> %#v", err)
	}

	log.Debugf("Send Mail ==> %#v", sendmail)

	auditLogResult, _ := repository.InsertAuditLog(model.AuditLog{
		MerchantId:      merchant["MerchantID"],
		UserId:          merchant["MerchantUID"],
		ActionName:      "CREATE",
		EntityType:      "TempMerchants",
		EntityId:        "",
		PropName:        "Status",
		OldValue:        "",
		NewValue:        "0",
		Remark:          "Approve",
		MenuName:        "Merchant Approve",
		ClientIpAddress: ipAddress,
		CreatedAt:       time.Now().Format("2006-01-02 15:04:05"),
	})
	log.Info("auditLogResult", auditLogResult)

	return "Register Success", nil
}

func (obj *mrService) DenyApproveMerchantRegister(data map[string]string, ipAddress string) (string, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	log.Debugf("data ==> %#v ", data)

	// check input
	if len(data["tempid"]) <= 0 || len(data["company"]) <= 0 || len(data["status"]) <= 0 {
		log.Errorf("Error : Incorrect information")
		return "Error", errors.New("Incorrect information")
	}

	// get TempsMerchants
	tempmerchant, err := repository.GetTempMerchant(data["tempid"], data["company"])
	if err != nil {
		log.Errorf("DB Error : %#v", err)
		return "Error", err
	}

	if data["tempid"] != tempmerchant.TempID {
		log.Errorf("Error : Incorrect information")
		return "Error", errors.New("Incorrect information")
	}

	log.Debugf("tempmerchant ==> %#v", tempmerchant)

	message := "-"
	if len(data["message"]) > 0 {
		message = data["message"]
	}

	username := "-"
	if len(data["username"]) > 0 {
		message = data["username"]
	}

	// update tempsmerchants
	uptemp, err := repository.UpdateDenyTempMerchants(data["tempid"], data["company"], 2, message, username)
	if err != nil {
		log.Errorf("DB Error : %#v", err)
	}

	log.Debugf("Update TempMerchant ==> %#v", uptemp)

	env := viper.GetString("SERVICE_ENV")

	tempatemail := ""
	subjectmail := ""
	if data["status"] == "1" {
		tempatemail = viper.GetString("TEMPLATE_EMAIL_DENY_NOUPDATE")
		if env != "Production" {
			subjectmail = "[Non-Prod] ร้านค้าไม่ผ่านเกณฑ์ เนื่องจากเว็บไซต์หรือช่องทางบริการไม่มีความเคลื่อนไหว | Inactivity on Website or Service Channels"
		} else {
			subjectmail = "ร้านค้าไม่ผ่านเกณฑ์ เนื่องจากเว็บไซต์หรือช่องทางบริการไม่มีความเคลื่อนไหว | Inactivity on Website or Service Channels"
		}
	} else if data["status"] == "2" {
		tempatemail = viper.GetString("TEMPLATE_EMAIL_DENY_ONDITIONS")
		if env != "Production" {
			subjectmail = "[Non-Prod] ร้านค้าไม่ผ่านเกณฑ์ เนื่องจากธุรกิจไม่เป็นตามเงื่อนไขการให้บริการของบริษัท | Merchant Approval Denied Due to Non-Compliance with Service Terms"
		} else {
			subjectmail = "ร้านค้าไม่ผ่านเกณฑ์ เนื่องจากธุรกิจไม่เป็นตามเงื่อนไขการให้บริการของบริษัท | Merchant Approval Denied Due to Non-Compliance with Service Terms"
		}
	} else if data["status"] == "3" {
		tempatemail = viper.GetString("TEMPLATE_EMAIL_DENY_UNRELIABLE")
		if env != "Production" {
			subjectmail = "[Non-Prod] ร้านค้าไม่ผ่านเกณฑ์ เนื่องจากงบการเงินหรือสถานะประกอบของบริษัทไม่น่าเชื่อถือ | Merchant Application Declined Due to Unreliable Financial Status"
		} else {
			subjectmail = "ร้านค้าไม่ผ่านเกณฑ์ เนื่องจากงบการเงินหรือสถานะประกอบของบริษัทไม่น่าเชื่อถือ | Merchant Application Declined Due to Unreliable Financial Status"
		}
	}

	log.Infof("-= Send Mail =-")
	mail := make(map[string]string)
	mail["ToEmail"] = tempmerchant.Email
	mail["TempID"] = data["tempid"]
	mail["EmailSupport"] = viper.GetString("EMAIL_SUPPORT")
	mail["TemplateMail"] = tempatemail
	mail["Subject"] = subjectmail
	sendmail, err := obj.mrHandler.AWSSendMail(mail)
	if err != nil {
		log.Errorf("Error sendmail ==> %#v", err)
	}

	log.Debugf("Send Mail ==> %#v", sendmail)

	auditLogResult, _ := repository.InsertAuditLog(model.AuditLog{
		MerchantId:      "",
		UserId:          data["username"],
		ActionName:      "UPDATE",
		EntityType:      "TempMerchants",
		EntityId:        data["tempid"],
		PropName:        "Register",
		OldValue:        "0",
		NewValue:        "2",
		Remark:          "Deny",
		MenuName:        "Merchant Approve",
		ClientIpAddress: ipAddress,
		CreatedAt:       time.Now().Format("2006-01-02 15:04:05"),
	})
	log.Info("auditLogResult", auditLogResult)

	return "Deny Success", nil
}

func (obj mrService) CreateNewResellerRegister(registermodel model.NewResellerRegisterModel) (string, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	// check input
	if len(registermodel.Email) <= 0 || len(registermodel.Company) <= 0 || len(registermodel.Name) <= 0 || len(registermodel.Phone) <= 0 {
		log.Errorf("Error : Incorrect information")
		return "Error", errors.New("Incorrect information")
	}

	// insert temp reseller
	tempReseller, err := obj.mrHandler.CreateNewReseller(registermodel)
	if err != nil {
		log.Errorf("DB Error : %#v", err)
		return "Error", err
	}

	log.Debugf("create temp reseller ==> %#v", tempReseller)

	if tempReseller != "COMPLETE" {
		return "Error", errors.New("Missing invalid field")
	}
	return "Register Success", nil
}
