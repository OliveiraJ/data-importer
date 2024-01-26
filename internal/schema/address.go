package schema

type Address struct {
	Street       string `csv:"Rua"`
	Number       string `csv:"Número"`
	Complement   string `csv:"Complemento"`
	Neighborhood string `csv:"Bairro"`
	CEP          string `csv:"CEP"`
	City         string `csv:"Cidade"`
	State        string `csv:"Estado (UF)"`
	Country      string `csv:"País"`
}

type MailingAddress struct {
	MailingStreet       string `csv:"Rua de correspondênci"`
	MailingNumber       string `csv:"Número de correspondência"`
	MailingComplement   string `csv:"Complemento de correspondência"`
	MailingNeighborhood string `csv:"Bairro de correspondência"`
	MailingCEP          string `csv:"CEP de correspondência"`
	MailingCity         string `csv:"Cidade de correspondência"`
	MailingState        string `csv:"Estado (UF) de correspondência"`
}
