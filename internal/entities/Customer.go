package entities

type Customer struct {
	DueDate           int
	CompanyName       string
	FantasyName       string
	Pronunciation     string
	BirthDate         string
	CNPJ              string
	CPF               string
	RG                string
	IssuingAuthority  string
	FinanceEmail      []string
	NotificationEmail []string
	Phone             string
	Cellphone         []string
	Street            string
	Number            int
	Complement        string
	Neighborhood      string
	City              string
	State             string
	ZipCode           string
}
