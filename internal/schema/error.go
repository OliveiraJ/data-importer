package schema

var ErrorKeys = map[string]int{
	"EMPTY_DUE_DATE":                 0,
	"EMPTY_COMPANY_NAME":             1,
	"INVALID_NOTIFICATION_EMAIL":     2,
	"INVALID_FINANCE_EMAIL":          3,
	"INVALID_CELLPHONE_NUMBER":       4,
	"INVALID_PRONE_NUMBER":           5,
	"INVALID_CPF_NUMBER":             6,
	"INVALID_CNPJ_NUMBER":            7,
	"INVALID_RG_NUMBER":              8,
	"CITY_NOT_FOUND":                 9,
	"UNKNOUNW_STATE":                 10,
	"INVALID_FOUNDATION_DATE_FORMAT": 11,
	"INVALID_REGISTER_DATE_FORMAT":   12,
	"INVALID_DUE_DATE":               13,
	"INVALID_CIVIL_STATE":            14,
	"REPEATED_REGISTER":              15,
}

type Error struct {
	Field       string
	Err         error
	Description string
}
