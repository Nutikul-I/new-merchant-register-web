package model

type (
	TokenResponse struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	CustomerRequest struct {
		CustomerName   string        `json:"customerName"`   // ชื่อลูกค้า Required [ชื่อลูกค้าห้ามซ้ำกัน]
		CustomerState  int           `json:"customerState"`  // <integer> สถานะของลูกค้า Required : 1 = Lead , 2 = Prospect , 3 = Customer
		CustomerStatus int           `json:"customerStatus"` // <integer> สถานะการใช้งาน 0 = active , 1 = ended (Required Note) , 2 = closed (Required Note)
		CustomerType   int           `json:"customerType"`   // <integer> ประเภท 1 = Business , 2 = individual
		SourceName     string        `json:"sourceName"`     // แหล่งที่มาลูกค้าถ้ามาจากเว็บใส่ Prefix: Z-{หมวดหมู่ธุรกิจ}
		LeadStatus     int           `json:"leadStatus"`     // สถานะลูกค้า 1 = New , 2 = Followup , 3 = Unqualified , 4 = Interested
		InterestsName  []string      `json:"interestsName"`  // ความสนใจ
		CustomFields   []CustomField `json:"customFields"`   // ข้อมูลฟิลด์เพิ่มเติม
		Contacts       []Contact     `json:"contacts"`       // ข้อมูลผู้ติดต่อ
		Owners         []string      `json:"owners"`         // รหัสพนังงาน
	}

	CustomField struct {
		CustomFieldName  string   `json:"customFieldName"`  // ชื่อหัวข้อฟิลด์ Required
		CustomFieldValue []string `json:"customFieldValue"` // ข้อมูลของฟิลด์ Required
	}

	Contact struct {
		ContactName   string `json:"contactName"`   // ผู้ติดต่อ Required
		ContactStatus bool   `json:"contactStatus"` // true = active , false = inactive
		ContactPhone  string `json:"contactPhone"`  // เบอร์โทรศัพท์ ของผู้ติดต่อ
		ContactMobile string `json:"contactMobile"` // เบอร์โทรศัพมือถือ ของผู้ติดต่อ
		ContactEmail  string `json:"contactEmail"`  // Email ของผู้ติดต่อ
	}
)
