package repository

import (
	"context"
	"database/sql"
	"register-service/model"
	"runtime"
	"strconv"
	"strings"

	"github.com/blockloop/scan"
	log "github.com/sirupsen/logrus"
)

func CheckMerchantCompany(company string) (bool, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	log.Debugf("company ==> %#v", company)

	db := GetDb()
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)

	if err != nil {
		log.Printf("%v", err)
		return false, err
	}

	// Execute query
	rows, err := db.QueryContext(ctx, model.SQL_Check_Merchant_Company, sql.Named("MerchantCompany", company))

	if err != nil {
		log.Printf(" %#v", err)
		return false, err
	}

	defer rows.Close()

	var count_merchant int

	err = scan.Row(&count_merchant, rows)

	if err != nil {
		log.Errorf(" %#v", err)
		return false, err

	}

	// log.Debugf("count company ==> %#v", count_merchant)

	if count_merchant == 1 {
		return false, nil
	}

	return true, nil
}

func CreateTempMerchant(merchant map[string]string) (string, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	// log.Debugf("merchant ==> %#v", merchant)

	db := GetDb()
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)

	if err != nil {
		log.Errorf("#%v", err)
		return "Error", err
	}

	stmt, err := db.Prepare(model.SQL_Add_Temp_Merchant)

	if err != nil {
		log.Errorf(" %#v", err)
		return "Error", err
	}
	_, err = stmt.Exec(
		merchant["ServiceTypeCode"],
		merchant["LangCode"],
		merchant["Name"],
		merchant["Surname"],
		merchant["Company"],
		merchant["Address"],
		merchant["Address2"],
		merchant["City"],
		merchant["State"],
		merchant["ProvinceID"],
		merchant["Postalcode"],
		merchant["CountryCode"],
		merchant["Tel"],
		merchant["Mobile"],
		merchant["Fax"],
		merchant["Email"],
		merchant["Website"],
		merchant["CatID"],
		merchant["SCatID"],
		merchant["SoSimple"],
		0,
		merchant["Agency"],
		"-",
		merchant["LeadNo"],
	)

	if err != nil {
		log.Errorf(" %#v", err)
		return "Error", err
	}

	return "COMPLETE", nil
}

func GetTempMerchantId(company string) (int, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	// log.Debugf("company ==> %#v", company)

	db := GetDb()
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)

	if err != nil {
		log.Printf("%v", err)
		return 0, err
	}

	// Execute query
	rows, err := db.QueryContext(ctx, model.SQL_Get_Temp_Merchant_Id, sql.Named("MerchantCompany", company))

	if err != nil {
		log.Printf(" %#v", err)
		return 0, err
	}

	defer rows.Close()

	var count_merchant int

	err = scan.Row(&count_merchant, rows)

	if err != nil {
		log.Errorf(" %#v", err)
		return 0, err

	}

	// log.Debugf("count company ==> %#v", count_merchant)

	return count_merchant, nil
}

func CreateTempBankMerchant(merchant map[string]string) (string, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	// log.Debugf("merchant ==> %#v", merchant)

	db := GetDb()
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)

	if err != nil {
		log.Errorf("#%v", err)
		return "Error", err
	}

	stmt, err := db.Prepare(model.SQL_Add_Temp_Bank_Merchant)

	if err != nil {
		log.Errorf(" %#v", err)
		return "Error", err
	}
	_, err = stmt.Exec(
		merchant["TempID"],
		merchant["BankCode"],
		merchant["AccountNo"],
		merchant["AccountName"],
		merchant["Branch"],
		merchant["AccountTypeCode"],
		merchant["UpCountryStatus"],
	)

	if err != nil {
		log.Errorf(" %#v", err)
		return "Error", err
	}

	return "COMPLETE", nil
}

func CreateMerchantCounterService(merchant map[string]string) (string, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	// log.Debugf("merchant ==> %#v", merchant)

	db := GetDb()
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)

	if err != nil {
		log.Errorf("#%v", err)
		return "Error", err
	}

	stmt, err := db.Prepare(model.SQL_Add_Merchant_CounterService)

	if err != nil {
		log.Errorf(" %#v", err)
		return "Error", err
	}
	_, err = stmt.Exec(
		merchant["MerchantID"],
		merchant["StartDate"],
		merchant["ExpiredDate"],
		"1",
		merchant["Enabled"],
		merchant["ServiceTypeCodeCS"],
	)

	if err != nil {
		log.Errorf(" %#v", err)
		return "Error", err
	}

	return "COMPLETE", nil
}

func CreateMerchantReseller(merchant map[string]string) (string, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	// log.Debugf("merchant ==> %#v", merchant)

	db := GetDb()
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)

	if err != nil {
		log.Errorf("#%v", err)
		return "Error", err
	}

	stmt, err := db.Prepare(model.SQL_Add_Merchant_Reseller)

	if err != nil {
		log.Errorf(" %#v", err)
		return "Error", err
	}
	_, err = stmt.Exec(
		merchant["Agency"],
		merchant["MerchantID"],
	)

	if err != nil {
		log.Errorf(" %#v", err)
		return "Error", err
	}

	return "COMPLETE", nil
}

func GetMerchantID(merchant map[string]string) (string, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	db := GetDb()
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)

	if err != nil {
		log.Printf("%v", err)
		return "0", err
	}

	// Execute query
	rows, err := db.QueryContext(ctx, model.SQL_Get_Merchant_ID, sql.Named("Password", merchant["Password"]), sql.Named("MerchantCompany", merchant["Company"]), sql.Named("MerchantEmail", merchant["Email"]))

	if err != nil {
		log.Printf(" %#v", err)
		return "0", err
	}

	defer rows.Close()

	var tranId string

	err = scan.Row(&tranId, rows)

	if err != nil {
		log.Errorf(" %#v", err)
		return "0", err

	}

	return tranId, nil
}

func CreateMerchant(merchant map[string]string) (string, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	log.Debugf("merchant ==> %#v", merchant)

	db := GetDb()
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)

	if err != nil {
		log.Errorf("#%v", err)
		return "Error", err
	}

	stmt, err := db.Prepare(model.SQL_Add_Merchant)

	if err != nil {
		log.Errorf(" %#v", err)
		return "Error", err
	}
	_, err = stmt.Exec(
		merchant["Password"],
		merchant["LangCode"],
		merchant["Name"],
		merchant["Name"],
		merchant["Name"],
		"-",
		merchant["Surname"],
		"-",
		merchant["Company"],
		merchant["Company"],
		merchant["Company"],
		"-",
		merchant["Address"],
		"-",
		"-",
		merchant["Address2"],
		"-",
		"-",
		merchant["City"],
		"-",
		0,
		merchant["ProvinceID"],
		0,
		"-",
		"-",
		merchant["Email"],
		merchant["Email"],
		"-",
		"-",
		merchant["Tel"],
		"-",
		merchant["Mobile"],
		"-",
		"-",
		merchant["Fax"],
		"-",
		"-",
		merchant["CountryCode"],
		"-",
		"-",
		merchant["Postalcode"],
		"-",
		merchant["StartDate"],
		merchant["ExpiredDate"],
		merchant["CurrencyCode"],
		merchant["Website"],
		merchant["Website"],
		"-",
		"-",
		"-",
		"Payment",
		merchant["ServiceTypeCode"],
		merchant["MonthlyFee"],
		merchant["VisaAccepted"],
		merchant["MasterCardAccepted"],
		merchant["JCBAccepted"],
		merchant["AMEXAccepted"],
		merchant["ChargePercent"],
		merchant["AMEXChargePercent"],
		merchant["BillingAddressEnable"],
		merchant["DemoStatus"],
		0,
		merchant["AutoRenew"],
		merchant["AccountEnable"],
		0,
		merchant["TransferMoneyPeriodCode"],
		"-",
		"DefaultTheme",
		"DefaultPage.master",
		"-",
		1,
		0.0000,
		0.0000,
		0.0000,
		0.0000,
		"-",
		merchant["CatID"],
		merchant["SCatID"],
		0,
		"-",
		0,
		0,
		0,
		0,
		"-",
		0,
		"-",
		"-",
		"-",
		"-",
		1,
		merchant["StartDate"],
		0,
		"-",
		0,
		0,
		4.75,
		0,
		0,
		0,
		3.5,
		0,
		0,
		0,
		3.75,
		3.75,
		3.75,
		3.75,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		3.75,
		3.75,
		3.75,
		3.75,
		3.75,
		"-",
		"-",
		0,
		"-",
		0,
		0.55,
		2,
		2,
		0,
		0,
		0,
		"P",
		0,
		0,
		0,
		0,
		0,
		merchant["LeadNo"],
	)

	if err != nil {
		log.Errorf(" %#v", err)
		return "Error", err
	}

	return "COMPLETE", nil
}

func UpdateMerchantUID(merchant map[string]string) (string, error) {
	/* Init log */
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})
	log.Debugf("merchant ==> %#v", merchant)

	db := GetDb()
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)

	if err != nil {
		log.Errorf("#%v", err)
		return "ERROR", err
	}

	stmt, err := db.Prepare(model.SQL_UPDATE_MERCHANTUID)

	if err != nil {
		log.Errorf(" %#v", err)
		return "ERROR", err
	}

	_, err = stmt.Exec(merchant["MerchantUID"],
		merchant["MerchantID"],
		merchant["Company"],
	)

	if err != nil {
		log.Errorf(" %#v", err)
		return "ERROR", err
	}

	return "COMPLETE", nil
}

func CreateChannelMerchant(merchant map[string]string, channel model.ChannelMerchant) (string, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	log.Debugf("channel ==> %#v", channel)

	db := GetDb()
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)

	if err != nil {
		log.Errorf("#%v", err)
		return "Error", err
	}

	stmt, err := db.Prepare(model.SQL_Add_Channel_Merchant)

	if err != nil {
		log.Errorf(" %#v", err)
		return "Error", err
	}
	_, err = stmt.Exec(merchant["MerchantID"],
		channel.ChannelCode,
		channel.ChannelType,
		"system",
		"system",
		0,
	)

	if err != nil {
		log.Errorf(" %#v", err)
		return "Error", err
	}

	return "COMPLETE", nil
}

func CreateMerchantConfiguration(merchant map[string]string) (string, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	log.Debugf("merchant ==> %#v", merchant)

	db := GetDb()
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)

	if err != nil {
		log.Errorf("#%v", err)
		return "Error", err
	}

	stmt, err := db.Prepare(model.SQL_Add_Merchant_Configuration)

	if err != nil {
		log.Errorf(" %#v", err)
		return "Error", err
	}
	_, err = stmt.Exec(merchant["MerchantID"],
		1,
		"-",
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		1,
	)

	if err != nil {
		log.Errorf(" %#v", err)
		return "Error", err
	}

	return "COMPLETE", nil
}

func GetUser(merchant map[string]string) (string, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	db := GetDb()
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)

	if err != nil {
		log.Printf("%v", err)
		return "0", err
	}

	// Execute query
	rows, err := db.QueryContext(ctx, model.SQL_Get_User, sql.Named("MerchantID", merchant["MerchantID"]))

	if err != nil {
		log.Printf(" %#v", err)
		return "0", err
	}

	defer rows.Close()

	var user_id string

	err = scan.Row(&user_id, rows)

	if err != nil {
		log.Errorf(" %#v", err)
		return "0", err

	}

	return user_id, nil
}

func CreateRole(user map[string]string) (string, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	log.Debugf("user ==> %#v", user)

	db := GetDb()
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)

	if err != nil {
		log.Errorf("#%v", err)
		return "Error", err
	}

	stmt, err := db.Prepare(model.SQL_Add_User_Role)

	if err != nil {
		log.Errorf(" %#v", err)
		return "Error", err
	}
	_, err = stmt.Exec(user["UserID"],
		user["UserRole"],
	)

	if err != nil {
		log.Errorf(" %#v", err)
		return "Error", err
	}

	return "COMPLETE", nil
}

func CountTempMerchantRegister(data map[string]string) (int, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	// log.Debugf("data ==> %#v", data)

	db := GetDb()
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)

	if err != nil {
		log.Printf("%v", err)
		return 0, err
	}
	query := model.SQL_Count_Temp_Merchant + " WHERE 1=1 "

	tempid := ""
	if len(data["tempid"]) > 1 {
		tempid = data["tempid"]
		query = query + " and TempID = " + tempid
	}

	email := ""
	if len(data["email"]) > 1 {
		email = data["email"]
		query = query + " and Email = '" + email + "'"
	}

	name := ""
	if len(data["name"]) > 1 {
		name = data["name"]
		query = query + " and Name = '" + name + "'"
	}

	// Execute query
	rows, err := db.QueryContext(ctx, query, sql.Named("StartDate", data["startdate"]), sql.Named("EndDate", data["enddate"]))

	if err != nil {
		log.Printf(" %#v", err)
		return 0, err
	}

	defer rows.Close()

	var count int

	err = scan.Row(&count, rows)

	if err != nil {
		log.Errorf(" %#v", err)
		return 0, err

	}

	// log.Debugf("count company ==> %#v", count_merchant)

	return count, nil
}

func GetTempMerchantRegister(data map[string]string) (model.TempMerchant, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	// log.Debugf("data ==> %#v", data)
	query := model.SQL_Get_Temp_Merchant + " WHERE 1=1 "

	tempid := ""
	if len(data["tempid"]) > 1 {
		tempid = data["tempid"]
		query = query + " and TempID = " + tempid
	}

	email := ""
	if len(data["email"]) > 1 {
		email = data["email"]
		query = query + " and Email = '" + email + "'"
	}

	name := ""
	if len(data["name"]) > 1 {
		name = data["name"]
		query = query + " and Name = '" + name + "'"
	}

	// order by f.TempID desc
	// OFFSET @Offset ROWS FETCH FIRST @Limit ROWS ONLY

	var offset = 0
	var limit = 20
	if data["limit"] != "" {
		limitint, _ := strconv.Atoi(data["limit"])
		limit = limitint
	}

	page, _ := strconv.Atoi(data["page"])

	if data["page"] == "1" {
		offset = 0
	} else {
		limitint, _ := strconv.Atoi(data["limit"])
		offset = limitint * (page - 1)
	}

	query = query + " order by f.TempID desc OFFSET " + strconv.Itoa(offset) + " ROWS FETCH FIRST " + strconv.Itoa(limit) + " ROWS ONLY"

	// log.Debugf("offset ==> %#v", offset)
	// log.Debugf("limit ==> %#v", limit)

	db := GetDb()
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)

	if err != nil {
		log.Printf("%v", err)
		return model.TempMerchant{}, err
	}

	// Execute query
	rows, err := db.QueryContext(ctx, query, sql.Named("StartDate", data["startdate"]), sql.Named("EndDate", data["enddate"]))

	if err != nil {
		log.Printf(" %#v", err)
		return model.TempMerchant{}, err
	}

	defer rows.Close()

	var TempMerchantData []model.TempMerchantData

	err = scan.Rows(&TempMerchantData, rows)

	if err != nil {
		log.Errorf(" %#v", err)
		return model.TempMerchant{}, err

	}

	var TempMerchant model.TempMerchant
	TempMerchant.Data = append(TempMerchant.Data, TempMerchantData...)

	// log.Debugf("count company ==> %#v", count_merchant)

	return TempMerchant, nil
}

func GetTempMerchant(tempid string, company string) (model.TempsMerchant, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	// log.Debugf("company ==> %#v", company)

	db := GetDb()
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)

	if err != nil {
		log.Printf("%v", err)
		return model.TempsMerchant{}, err
	}

	// Execute query
	rows, err := db.QueryContext(ctx, model.SQL_Get_Temps_Merchant, sql.Named("TempID", tempid), sql.Named("Company", company))

	if err != nil {
		log.Printf(" %#v", err)
		return model.TempsMerchant{}, err
	}

	defer rows.Close()

	var count_merchant model.TempsMerchant

	err = scan.Row(&count_merchant, rows)

	if err != nil {
		log.Errorf(" %#v", err)
		return model.TempsMerchant{}, err

	}

	// log.Debugf("count company ==> %#v", count_merchant)

	return count_merchant, nil
}

func UpdateTempMerchants(tempid string, company string, message string, merchantuid string, username string) (string, error) {
	/* Init log */
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})
	log.Debugf("tempid ==> %#v", tempid)
	log.Debugf("company ==> %#v", company)
	log.Debugf("message ==> %#v", message)
	log.Debugf("merchantuid ==> %#v", merchantuid)
	log.Debugf("username ==> %#v", username)

	db := GetDb()
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)

	if err != nil {
		log.Errorf("#%v", err)
		return "ERROR", err
	}

	stmt, err := db.Prepare(model.SQL_UPDATE_TMEP_MERCHANT)

	if err != nil {
		log.Errorf(" %#v", err)
		return "ERROR", err
	}

	_, err = stmt.Exec(1,
		tempid,
		company,
		message,
		merchantuid,
		username,
	)

	if err != nil {
		log.Errorf(" %#v", err)
		return "ERROR", err
	}

	return "COMPLETE", nil
}

func UpdateDenyTempMerchants(tempid string, company string, status int, message string, username string) (string, error) {
	/* Init log */
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})
	log.Debugf("tempid ==> %#v", tempid)
	log.Debugf("company ==> %#v", company)
	log.Debugf("status ==> %#v", status)
	log.Debugf("message ==> %#v", message)
	log.Debugf("username ==> %#v", username)

	db := GetDb()
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)

	if err != nil {
		log.Errorf("#%v", err)
		return "ERROR", err
	}

	log.Debugf("sql ==> %#v", model.SQL_UPDATE_DENY_TMEP_MERCHANT)

	stmt, err := db.Prepare(model.SQL_UPDATE_DENY_TMEP_MERCHANT)

	if err != nil {
		log.Errorf(" %#v", err)
		return "ERROR", err
	}

	_, err = stmt.Exec(
		tempid,
		status,
		message,
		company,
		username,
	)

	if err != nil {
		log.Errorf(" %#v", err)
		return "ERROR", err
	}

	return "COMPLETE", nil
}
