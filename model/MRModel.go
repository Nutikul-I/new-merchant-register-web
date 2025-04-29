package model

type ResponseApprove struct {
	Status  string `json:"Status"`
	Message string `json:"Message"`
}

type TempMerchant struct {
	TotalPage int `json:"TotalPage"`
	Page      int `json:"Page"`
	NextPage  int `json:"NextPage"`
	TotalData int `json:"TotalData"`
	Data      []TempMerchantData
}

type TempMerchantData struct {
	TempID          int    `json:"TempID"`
	MerchantUID     string `json:"MerchantUID"`
	ServiceTypeCode string `json:"ServiceTypeCode"`
	LangCode        string `json:"LangCode"`
	Name            string `json:"Name"`
	Surname         string `json:"Surname"`
	Company         string `json:"Company"`
	Address         string `json:"Address"`
	Address2        string `json:"Address2"`
	City            string `json:"City"`
	State           string `json:"State"`
	Province        string `json:"Province"`
	Postalcode      string `json:"Postalcode"`
	CountryCode     string `json:"CountryCode"`
	Mobile          string `json:"Mobile"`
	Email           string `json:"Email"`
	Website         string `json:"Website"`
	Catagory        string `json:"Catagory"`
	SubCatgory      string `json:"SubCatgory"`
	CreateDateTime  string `json:"CreateDateTime"`
	Status          string `json:"Status"`
	Username        string `json:"ByUsername"`
	UpdateDateTime  string `json:"UpdateDateTime"`
	Message         string `json:"Message"`
	ApproveBy       string `json:"ApproveBy"`
	LeadNo          string `json:"LeadNo"`
}

type TempsMerchant struct {
	TempID         string `json:"TempID"`
	Name           string `json:"Name"`
	Surname        string `json:"Surname"`
	Company        string `json:"Company"`
	LangCode       string `json:"LangCode"`
	Mobile         string `json:"Mobile"`
	Email          string `json:"Email"`
	Website        string `json:"Website"`
	CatID          string `json:"CatID"`
	SCatID         string `json:"SCatID"`
	CreateDateTime string `json:"CreateDateTime"`
	Register       string `json:"Register"`
	Agency         string `json:"Agency"`
	Username       string `json:"Username"`
	LeadNo         string `json:"LeadNo"`
}

type RegisterModel struct {
	MerchantID              string  `json:"merchantid"`
	MerchantUID             string  `json:"merchantuid"`
	Name                    string  `json:"name"`
	Mobile                  string  `json:"mobile"`
	Email                   string  `json:"email"`
	Category                string  `json:"category"`
	CategoryNameTH          string  `json:"categorynameTH"`
	CategoryNameEN          string  `json:"categorynameEN"`
	Subcategory             string  `json:"subcategory"`
	LeadNo                  string  `json:"leadno"`
	Website                 string  `json:"website"`
	Company                 string  `json:"company"`
	Check                   string  `json:"check"`
	Lang                    string  `json:"lang"`
	Level                   string  `json:"level"`
	ServiceTypeCode         string  `json:"serviceTypeCode"`
	LangCode                string  `json:"langCode"`
	Surname                 string  `json:"surname"`
	Address                 string  `json:"address"`
	Address2                string  `json:"address2"`
	City                    string  `json:"city"`
	State                   string  `json:"state"`
	ProvinceID              int     `json:"provinceID"`
	Postalcode              string  `json:"postalcode"`
	CountryCode             string  `json:"countryCode"`
	Tel                     string  `json:"tel"`
	Fax                     string  `json:"fax"`
	SoSimple                int     `json:"soSimple"`
	Password                string  `json:"password"`
	StartDate               string  `json:"startDate"`
	ExpiredDate             string  `json:"expiredDate"`
	CurrencyCode            string  `json:"currencyCode"`
	MonthlyFee              int     `json:"monthlyFee"`
	VisaAccepted            int     `json:"visaAccepted"`
	MasterCardAccepted      int     `json:"masterCardAccepted"`
	AMEXAccepted            int     `json:"aMEXAccepted"`
	ChargePercent           float64 `json:"chargePercent"`
	AMEXChargePercent       float64 `json:"aMEXChargePercent"`
	BillingAddressEnable    int     `json:"billingAddressEnable"`
	DemoStatus              int     `json:"demoStatus"`
	AutoRenew               int     `json:"autoRenew"`
	AccountEnable           int     `json:"accountEnable"`
	TransferMoneyPeriodCode string  `json:"transferMoneyPeriodCode"`
	HideMenuPaysbuy         int     `json:"hideMenuPaysbuy"`
	Enabled                 int     `json:"enabled"`
	ServiceTypeCodeCS       string  `json:"serviceTypeCodeCS"`
	LevelPaysocial          string  `json:"levelPaysocial"`
	BankCode                string  `json:"bankcode"`
	AccountNo               string  `json:"AccountNo"`
	AccountName             string  `json:"AccountName"`
	Branch                  string  `json:"Branch"`
	Accounttype             string  `json:"AccountType"`
	Upcountry               string  `json:"upcountry"`
	Agency                  string  `json:"Agency"`
}

type NewRegisterModel struct {
	MerchantID              string  `json:"merchantid"`
	MerchantUID             string  `json:"merchantuid"`
	Name                    string  `json:"name"`
	Mobile                  string  `json:"mobile"`
	Email                   string  `json:"email"`
	Category                string  `json:"category"`
	CategoryNameTH          string  `json:"categorynameTH"`
	CategoryNameEN          string  `json:"categorynameEN"`
	Subcategory             string  `json:"subcategory"`
	Website                 string  `json:"website"`
	Company                 string  `json:"company"`
	Check                   string  `json:"check"`
	Lang                    string  `json:"lang"`
	Level                   string  `json:"level"`
	ServiceTypeCode         string  `json:"serviceTypeCode"`
	LangCode                string  `json:"langCode"`
	Surname                 string  `json:"surname"`
	Address                 string  `json:"address"`
	Address2                string  `json:"address2"`
	City                    string  `json:"city"`
	State                   string  `json:"state"`
	ProvinceID              int     `json:"provinceID"`
	Postalcode              string  `json:"postalcode"`
	CountryCode             string  `json:"countryCode"`
	Tel                     string  `json:"tel"`
	Fax                     string  `json:"fax"`
	SoSimple                int     `json:"soSimple"`
	Password                string  `json:"password"`
	StartDate               string  `json:"startDate"`
	ExpiredDate             string  `json:"expiredDate"`
	CurrencyCode            string  `json:"currencyCode"`
	MonthlyFee              int     `json:"monthlyFee"`
	VisaAccepted            int     `json:"visaAccepted"`
	MasterCardAccepted      int     `json:"masterCardAccepted"`
	AMEXAccepted            int     `json:"aMEXAccepted"`
	ChargePercent           float64 `json:"chargePercent"`
	AMEXChargePercent       float64 `json:"aMEXChargePercent"`
	BillingAddressEnable    int     `json:"billingAddressEnable"`
	DemoStatus              int     `json:"demoStatus"`
	AutoRenew               int     `json:"autoRenew"`
	AccountEnable           int     `json:"accountEnable"`
	TransferMoneyPeriodCode string  `json:"transferMoneyPeriodCode"`
	HideMenuPaysbuy         int     `json:"hideMenuPaysbuy"`
	Enabled                 int     `json:"enabled"`
	ServiceTypeCodeCS       string  `json:"serviceTypeCodeCS"`
	LevelPaysocial          string  `json:"levelPaysocial"`
	BankCode                string  `json:"bankcode"`
	AccountNo               string  `json:"AccountNo"`
	AccountName             string  `json:"AccountName"`
	Branch                  string  `json:"Branch"`
	Accounttype             string  `json:"AccountType"`
	Upcountry               string  `json:"upcountry"`
	Agency                  string  `json:"Agency"`
	LeadSource              string  `json:"leadsource"`
	CF1129                  string  `json:"cf1129"`
	CF1148                  string  `json:"cf1148"`
	CF1150                  string  `json:"cf1150"`
	CF1152                  string  `json:"cf1152"`
	CF1154                  string  `json:"cf1154"`
	CF1156                  string  `json:"cf1156"`
}

type EncryptTripleDES struct {
	Status  int    `json:"Status"`
	Data    string `json:"Data"`
	Message string `json:"Message"`
}

type ChannelMerchant struct {
	ChannelCode string `json:"ChannelCode"`
	ChannelType string `json:"ChannelType"`
}

type User struct {
	CreateBy     string `json:"createBy"`
	UpdateBy     string `json:"updateBy"`
	CreateDate   string `json:"createDate"`
	UpdateDate   string `json:"updateDate"`
	EntityStatus bool   `json:"entityStatus"`
	UserName     string `json:"userName"`
	UserPass     string `json:"userPass"`
	UserDesc     string `json:"userDesc"`
	UserEmail    string `json:"userEmail"`
	MerchantId   string `json:"merchantId"`
	Cid          string `json:"cid"`
	Links        struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		AppUser struct {
			Href string `json:"href"`
		} `json:"appUser"`
		Roles struct {
			Href string `json:"href"`
		} `json:"roles"`
	} `json:"_links"`
}

type SecurePassword struct {
	Password string `json:"password"`
}

type SendMail struct {
	Message string `json:"message"`
}

type Category struct {
	MainValue int    `json:"mainvalue"`
	Value     int    `json:"value"`
	NameTH    string `json:"nameTH"`
	NameEN    string `json:"nameEN"`
}

type TurnstileResponse struct {
	Success     bool     `json:"success"`
	ChallengeTs string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes"`
}

var SQL_Check_Merchant_Company = `select SUM(TempID) as TempID from (
select count(TempID) as TempID from TempMerchants tm where tm.Company = @MerchantCompany
UNION ALL   
select count(MerchantID) as TempID from Merchant m where m.MerchantCompany = @MerchantCompany
) f`

var SQL_Add_Temp_Merchant = `INSERT INTO TempMerchants (ServiceTypeCode, LangCode, Name, Surname, Company, Address, Address2, City, State, ProvinceID, Postalcode, CountryCode, Tel, Mobile, Fax, Email, Website, CatID, SCatID, SosimpleFlag, CreateDateTime, Register, Agency, Message, UpdateDateTime, LeadNo)
	VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12, @p13, @p14, @p15, @p16, @p17, @p18, @p19, @p20, getdate(), @p21, @p22, @p23, getdate(), @p24)`

var SQL_Get_Temp_Merchant_Id = `select ISNULL(TempID,0) as TempID from TempMerchants where Company = @MerchantCompany order by TempID desc`

var SQL_Add_Temp_Bank_Merchant = `INSERT INTO TempMerchantsBankAccount (TempID, BankCode, AccountNo, AccountName, Branch, AccountTypeCode, UpCountryStatus)
	VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7);`

var SQL_Get_Merchant_ID = `select ISNULL(MerchantID,0) as MerchantID from Merchant where MerchantPassword = @Password and MerchantCompany = @MerchantCompany and MerchantEmail= @MerchantEmail;`

var SQL_Add_Merchant = `INSERT INTO Merchant
(MerchantPassword,
LangCode,
MerchantName,
NameTH,
NameEN,
MerchantSurname,
SurnameTH,
SurnameEN,
MerchantCompany,
CompanyTH,
CompanyEN,
MerchantAddress,
AddressTH,
AddressEN,
MerchantAddress2,
Address2TH,
Address2EN,
City,
CityTH,
CityEN,
ProvinceID,
ProvinceIDTH,
ProvinceIDEN,
ProvinceName,
Email,
OrdersEmail,
MerchantEmail,
MerchantContactEmail,
MerchantTel,
TelTH,
TelEN,
MerchantMobile,
Mobile,
MerchantFax,
FaxTH,
FaxEN,
MerchantCountryCode,
CountryCodeTH,
CountryCodeEN,
MerchantPostcode,
PostalcodeTH,
PostalcodeEN,
MerchantStartDate,
MerchantExpiredDate,
MerchantCurrencyCode,
DomainName,
MerchantURL,
MerchantReturnURL,
DetailTH,
DetailEN,
ServiceType,
ServiceTypeCode,
MonthlyFee,
VisaAccepted,
MasterCardAccepted,
JCBAccepted,
AMEXAccepted,
ChargePercent,
AmexChargePercent,
BillingAddressEnable,
DemoStatus,
CharityStatus,
AutoRenew,
AccountEnable,
ChangeCurrencyEnable,
TransferMoneyPeriodCode,
Logo,
Theme,
MasterPage,
Comment,
ShippingMethodEnable,
FreeShippingThreshold,
DiscountThreshold,
DiscountPercent,
ExtraChargePercent,
ExtraChargeDetail,
CatID,
SubCatID,
SMSAlertStatus,
SMSMobile,
SMSAlert30Days,
SMSAlert15Days,
SMSAlert7Days,
SMSAlertToday,
SMSBankDetail,
KbankMerchantEnable,
KbankMerchantID,
KbankTerminalID,
KbankMerchantName,
PriorityIndex,
CheckUpdateStep,
LastUpdate,
SosimpleFlag,
CVtype,
InstallmentFlag,
UnionpayAccepted,
PaypalChargePercent,
PaypalFlag,
BitcoinFlag,
mPOSAccepted,
mPOSChargePercent,
CODFlag,
ResellerID,
DirectFlag,
CODChargePercent,
CounterserviceChargePercent,
BillOtherChargePercent,
InternetBankingChargePercent,
AcceptedOnlyThaiIssuer,
KbankInstallmentAccepted,
BAYInstallmentAccepted,
KTCInstallmentAccepted,
BBLInstallmentAccepted,
SCBInstallmentAccepted,
BAYiBankingAccepted,
BBLiBankingAccepted,
SCBiBankingAccepted,
KTBiBankingAccepted,
AlipayAccepted,
WeChatPayAccepted,
BillPayAccepted,
VisaChargePercent,
MasterChargePercent,
CUPChargePercent,
JCBChargePercent,
InstallmentChargePercent,
RefundPolicyTH,
RefundPolicyEN,
KbankMerchant,
SecretKey,
PromptPayAccepted,
PromptPayChargePercent,
AlipayChargePercent,
WeChatPayChargePercent,
KbankiBankingAccepted,
RecurringAccepted,
TbankiBankingAccepted,
TypeRatePP,
TrueMoneyAccepted,
TokenizationAccepted,
CryptoAccepted,
XnapAccepted,
AtomeInstallmentAccepted,
LeadNo) VALUES (@p1,@p2,@p3,@p4,@p5,@p6,@p7,@p8,@p9,@p10,@p11,@p12,@p13,@p14,@p15,@p16,@p17,@p18,@p19,@p20,@p21,@p22,@p23,@p24,@p25,@p26,@p27,@p28,@p29,@p30,@p31,@p32,@p33,@p34,@p35,@p36,@p37,@p38,@p39,@p40,@p41,@p42,@p43,@p44,@p45,@p46,@p47,@p48,@p49,@p50,@p51,@p52,@p53,@p54,@p55,@p56,@p57,@p58,@p59,@p60,@p61,@p62,@p63,@p64,@p65,@p66,@p67,@p68,@p69,@p70,@p71,@p72,@p73,@p74,@p75,@p76,@p77,@p78,@p79,@p80,@p81,@p82,@p83,@p84,@p85,@p86,@p87,@p88,@p89,@p90,@p91,@p92,@p93,@p94,@p95,@p96,@p97,@p98,@p99,@p100,@p101,@p102,@p103,@p104,@p105,@p106,@p107,@p108,@p109,@p110,@p111,@p112,@p113,@p114,@p115,@p116,@p117,@p118,@p119,@p120,@p121,@p122,@p123,@p124,@p125,@p126,@p127,@p128,@p129,@p130,@p131,@p132,@p133,@p134,@p135,@p136,@p137,@p138,@p139,@p140,@p141,@p142,@p143,@p144);`

var SQL_UPDATE_MERCHANTUID = `UPDATE Merchant 
SET MerchantUID=@p1,LastUpdate=getdate()
WHERE MerchantID=@p2 and MerchantCompany = @p3;`

var SQL_Add_Merchant_CounterService = `INSERT INTO MerchantCounterService (MerchantID,StartDate,ExpireDate,HideMenuPaysbuy,Enabled,ServiceTypeCodeCS)
VALUES (@p1, @p2, @p3, @p4, @p5, @p6);`

var SQL_Add_Merchant_Reseller = `INSERT INTO MerchantWithReseller (ResellerID,MerchantID)
VALUES (@p1, @p2);`

var SQL_Add_Channel_Merchant = `INSERT INTO PaysolutionChannelMerchant (MerchantID ,ChannelCode ,ChannelType ,CreatedBy ,UpdatedBy ,CreatedDate ,UpdatedDate ,Status) 
VALUES (@p1, @p2, @p3, @p4, @p5, getdate(), getdate(), @p6)`

var SQL_Add_Merchant_Configuration = `INSERT INTO MerchantConfiguration (MerchantID ,ChangeAcceptedCard ,BriefMerchantDetail ,EnableClientList ,EnablePostBackUrl ,PostMerchantID ,PostCustomerEmail ,PostProductDetail ,PostTotal ,PostCardType ,PaymentPerMonth ,PaymentPerTran ,IsRateSetting ,IsNewPaymentSetting) 
VALUES (@p1,@p2,@p3,@p4,@p5,@p6,@p7,@p8,@p9,@p10,@p11,@p12,@p13,@p14)`

var SQL_Get_User = `select Top 1 id from app_user au where merchant_id = @MerchantID order by id desc`

var SQL_Add_User_Role = `INSERT INTO user_role (user_id,role_id)
VALUES (@p1, @p2);`

var SQL_Count_Temp_Merchant = `select Count(TempID) as Total from (select TempID ,ServiceTypeCode ,LangCode ,ISNULL(Name,'-') as Name ,ISNULL(Surname,'-') as Surname ,ISNULL(Company ,'-') as Company ,Address ,Address2 ,City ,State ,
ISNULL((select NameTH from Province p where p.ProvinceID = tm.ProvinceID),'-') as Province ,
Postalcode ,CountryCode ,Mobile ,Email ,Website ,
ISNULL((select CategoryNameTH from MerchantCategory mc where mc.CatID = tm.CatID  and mc.Enable = 1 ),'-') as Catagory ,
ISNULL((select SubCategoryNameTH from MerchantSubCategory msc where msc.CatID = tm.CatID and msc.SubCatID = tm.SCatID and msc.Enable = 1),'-') as SubCatgory ,
CreateDateTime , Register as Status
from TempMerchants tm
where tm.CreateDateTime >= @StartDate and tm.CreateDateTime <= @EndDate
) f`

var SQL_Get_Temp_Merchant = `select * from (
select TempID 
,ISNULL(MerchantUID,'-') as MerchantUID 
,ServiceTypeCode 
,LangCode 
,ISNULL(Name,'-') as Name 
,ISNULL(Surname,'-') as Surname 
,ISNULL(Company ,'-') as Company 
,Address 
,Address2 
,City 
,State 
,ISNULL((select NameTH from Province p where p.ProvinceID = tm.ProvinceID),'-') as Province 
,Postalcode ,CountryCode ,Mobile ,Email ,Website 
,ISNULL(Username,'-') as Username 
,ISNULL((select CategoryNameTH from MerchantCategory mc where mc.CatID = tm.CatID  and mc.Enable = 1 ),'-') as Catagory 
,ISNULL((select SubCategoryNameTH from MerchantSubCategory msc where msc.CatID = tm.CatID and msc.SubCatID = tm.SCatID and msc.Enable = 1),'-') as SubCatgory 
,CreateDateTime 
,ISNULL(UpdateDateTime,CreateDateTime) as UpdateDateTime 
,IIF(Register = 0 ,'wait approve',IIF(Register = 1 , 'approve', IIF(Register = 2 , 'not approved','wait approve'))) as Status 
,ISNULL(Message, '-') as Message
,ISNULL(Username , '-') as ApproveBy
,ISNULL(LeadNo , '-') as LeadNo
from TempMerchants tm
where tm.CreateDateTime >= @StartDate and tm.CreateDateTime <= @EndDate
) f`

var SQL_Get_Temps_Merchant = `select TOP 1 TempID, Name, Surname, Company, LangCode, Mobile, Email, Website, CatID, SCatID, CreateDateTime, Register, ISNULL(Agency,'-') as Agency, ISNULL(Username,'-') as Username, ISNULL(LeadNo,'-') as LeadNo from TempMerchants tm where TempID = @TempID and Company = @Company order by TempID desc`

var SQL_UPDATE_TMEP_MERCHANT = `UPDATE TempMerchants SET Register = @p1 ,Message = @p4 ,MerchantUID = @p5 ,Username = @p6 ,UpdateDateTime = getdate() WHERE TempID = @p2 and Company = @p3;`

var SQL_UPDATE_DENY_TMEP_MERCHANT = `UPDATE TempMerchants SET Register = @p2, Message = @p3, Username = @p5 ,UpdateDateTime = getdate()  WHERE TempID = @p1 and Company = @p4;`

type WithdrawAuto struct {
	Data struct {
		Transaction  string `json:"Transaction"`
		ResponseCode string `json:"ResponseCode"`
		Description  string `json:"Description"`
	} `json:"data"`
	Msg         string `json:"msg"`
	ApiVersion  string `json:"apiVersion"`
	CurrentDate string `json:"currentDate"`
}
