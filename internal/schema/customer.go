package schema

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"github.com/Nhanderu/brdoc"
)

type Customer struct {
	CustomerId           string `csv:"ID do Cliente"`
	DueDate              string `csv:"Dia de vencimento *"`
	CompanyName          string `csv:"Razão social ou Nome completo *"`
	TradeName            string `csv:"Nome fantasia"`
	Pronuntiation        string `csv:"Pronúncia"`
	FundationDate        string `csv:"Data de nascimento ou fundação"`
	CNPJ                 string `csv:"CNPJ"`
	CPF                  string `csv:"CPF"`
	RG                   string `csv:"RG"`
	IssuingAuthority     string `csv:"Órgão expedidor"`
	FinanceEmail         string `csv:"E-mail de financeiro"`
	NotificationEmail    string `csv:"E-mail de recado"`
	Telephone            string `csv:"Telefone"`
	Cellphone            string `csv:"Celular"`
	DDR                  string `csv:"DDR"`
	Extension            string `csv:"Ramal exclusivo"`
	Mailbox              string `csv:"Caixa postal"`
	Passport             string `csv:"Passaporte"`
	LateFee              string `csv:"Multa por atraso (%)"`
	DayFee               string `csv:"Juros por dia (%)"`
	DelayDaysForBlocking string `csv:"Dias de atraso para bloqueio"`
	ISSRetainer          string `csv:"Retentor ISS"`
	MunicipalInscription string `csv:"Inscrição municipal"`
	StateInscription     string `csv:"Inscrição estadual"`
	Site                 string `csv:"Site"`
	BranchOfActivity     string `csv:"Ramo de atividade"`
	Notes                string `csv:"Observações"`
	Presentation         string `csv:"Apresentação"`
	ServiceDescription   string `csv:"Descrição de produtos/serviços"`
	Treatment            string `csv:"Tratamento"`
	ServiceOrientation   string `csv:"Orientação para atendimento"`
	Occuptation          string `csv:"Profissão"`
	CivilState           string `csv:"Estado civil"`
	RegisterDate         string `csv:"Data do cadastro"`
	UnityID              string `csv:"ID da Unidade"`
	Address
	MailingAddress
	Errors map[int]Error
}

// Executa uma série de funções de validação, responsáveis por preencher possíveis erros nos campos
func (c *Customer) Validate(customers *[]Customer) {
	c.Errors = make(map[int]Error)
	c.validateRequiredFields()
	c.validateDueDate()
	c.validateEmails()
	c.validateDocuments()
	c.validatePhone()
	c.validateDate()
	c.validateCivilState()
	c.validateRepeatedCustomers(customers)

	for _, e := range c.Errors {
		if e.Err != nil {
			fmt.Println(c.CompanyName, e)
		}
	}
}

func (c *Customer) validateRequiredFields() {
	if c.DueDate == "" {
		c.Errors[ErrorKeys["EMPTY_DUE_DATE"]] = Error{
			Err: fmt.Errorf("%s", "O campo dia de vencimento não pode ser nulo"),
		}
	}

	if c.CompanyName == "" {
		c.Errors[ErrorKeys["EMPTY_COMPANY_NAME"]] = Error{
			Err: fmt.Errorf("%s", "O campo razão social não pode ser nulo"),
		}
	}
}

func (c *Customer) validateDueDate() {
	dueDate, err := strconv.Atoi(c.DueDate)
	if err != nil {
		panic(err)
	}

	if dueDate < 1 || dueDate > 28 {
		c.Errors[ErrorKeys["INVALID_DUE_DATE"]] = Error{
			Err: fmt.Errorf("%s", "O campo dia de vencimento tem que ter valor entre 1 e 28"),
		}
	}
}

func (c *Customer) validateEmails() {

	if c.NotificationEmail == "" && c.FinanceEmail == "" {
		return
	}

	notificationEmails := strings.Split(c.NotificationEmail, ";")
	financeEmails := strings.Split(c.FinanceEmail, ";")

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)

	for _, e := range notificationEmails {
		if !re.MatchString(e) {
			c.Errors[ErrorKeys["INVALID_NOTIFICATION_EMAIL"]] = Error{
				Err: fmt.Errorf("%s", "Emails de recados inválidos"),
			}
		}
	}

	for _, e := range financeEmails {
		if !re.MatchString(e) {
			c.Errors[ErrorKeys["INVALID_FINANCE_EMAIL"]] = Error{
				Err: fmt.Errorf("%s", "Emails financeiro inválidos"),
			}
		}
	}
}

func (c *Customer) validateDocuments() {
	if c.CPF != "" && !brdoc.IsCPF(c.CPF) {
		c.Errors[ErrorKeys["INVALID_CPF_NUMBER"]] = Error{
			Err: fmt.Errorf("%s", "CPF inválido"),
		}
		fmt.Println("CPF inválido")
	}

	if c.CNPJ != "" && !brdoc.IsCNPJ(c.CNPJ) {
		c.Errors[ErrorKeys["INVALID_CNPJ_NUMBER"]] = Error{
			Err: fmt.Errorf("%s", "CNPJ inválido"),
		}
		fmt.Println("CNPJ inválido")
	}
}

func (c *Customer) validatePhone() {
	if c.Cellphone != "" {

		cellPhoneRegex := `^\d{2}9\d{8}$`
		re := regexp.MustCompile(cellPhoneRegex)

		var cleanCellPhone []rune

		for _, char := range c.Cellphone {
			if unicode.IsDigit(char) {
				cleanCellPhone = append(cleanCellPhone, char)
			}
		}

		code := string(cleanCellPhone[:2])

		//firsDigit, _ := strconv.Atoi(string(cleanCellPhone[3])) -> verificar se essa condição é verdadeira

		if !re.MatchString(string(cleanCellPhone)) || !slices.Contains(Area_codes, code) {
			c.Errors[ErrorKeys["INVALID_CELLPHONE_NUMBER"]] = Error{
				Err: fmt.Errorf("%s", "Número de celular inválido"),
			}
		}

	}

	if c.Telephone != "" {
		telephoneRegex := `^\(?\d{2}\)?[-.\s]?\d{4,5}[-.\s]?\d{4}$`
		re := regexp.MustCompile(telephoneRegex)

		telephones := strings.Split(c.Telephone, ";")

		for _, t := range telephones {
			if !re.MatchString(t) {
				c.Errors[ErrorKeys["INVALID_PRONE_NUMBER"]] = Error{
					Err: fmt.Errorf("%s", "Número de telefone inválido"),
				}
			}
		}
	}
}

func (c *Customer) validateDate() {
	dateRegex := `^(0[1-9]|[12][0-9]|3[01])/(0[1-9]|1[0-2])/\d{4}$`

	re := regexp.MustCompile(dateRegex)

	if c.FundationDate != "" && !re.MatchString(c.FundationDate) {
		c.Errors[ErrorKeys["INVALID_FOUNDATION_DATE_FORMAT"]] = Error{
			Err: fmt.Errorf("%s", "Data de fundação/nascimento inválida"),
		}
	}

	if c.RegisterDate != "" && !re.MatchString(c.RegisterDate) {
		c.Errors[ErrorKeys["INVALID_REGISTER_DATE_FORMAT"]] = Error{
			Err: fmt.Errorf("%s", "Data de cadastro inválida"),
		}
	}
}

func (c *Customer) validateCivilState() {
	civilStates := []string{"C", "S", "D", "V"}
	if c.CPF != "" {
		if c.CivilState != "" && !slices.Contains(civilStates, c.CivilState) {
			c.Errors[ErrorKeys["INVALID_CIVIL_STATE"]] = Error{
				Err: fmt.Errorf("%s", "O estado civil informado é inválido"),
			}
		}
	}
}

func (c *Customer) validateRepeatedCustomers(customers *[]Customer) {
	count := 0
	for _, d := range *customers {
		if c.CompanyName == d.CompanyName {
			count++
		}
	}

	if count > 1 {
		c.Errors[ErrorKeys["REPEATED_REGISTER"]] = Error{
			Err: fmt.Errorf("%s", "Cliente repetido"),
		}
	}
}
