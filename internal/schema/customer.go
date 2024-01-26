package schema

import "fmt"

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
}

func (c Customer) Validate() error {
	return fmt.Errorf("%s", "erro")
}
