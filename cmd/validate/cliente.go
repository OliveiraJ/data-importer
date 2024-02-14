package validate

import (
	"github.com/OliveiraJ/data-importer/file"
	"github.com/OliveiraJ/data-importer/internal/schema"
	"github.com/spf13/cobra"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var report bool

func validateCliente() *cobra.Command {
	clienteCmd := &cobra.Command{
		Use:   "cliente",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {

			customers := []schema.Customer{}

			file.ReadFile(&customers)

			//fmt.Print(setupTable(customers))
			for _, v := range customers {
				v.Validate(&customers)
			}

		},
	}

	clienteCmd.Flags().BoolVarP(&report, "relatório", "r", false, "Se um relatório de consistência deve ser gerado")

	return clienteCmd
}

func setupTable(customers []schema.Customer) *table.Table {
	columns := []string{
		"ID do Cliente",
		"Dia de vencimento *",
		"Razão social ou Nome completo *",
		"Nome fantasia",
		"Pronúncia",
		"Data de nascimento ou fundação",
		"CNPJ",
		"CPF",
		"RG",
		"Órgão expedidor",
		"E-mail de financeiro",
		"E-mail de recado",
		"Telefone",
		"Celular",
		"DDR",
		"Ramal exclusivo",
		"Caixa postal",
		"Passaporte",
		"Multa por atraso (%)",
		"Juros por dia (%)",
		"Dias de atraso para bloqueio",
		"Retentor ISS",
		"Inscrição municipal",
		"Inscrição estadual",
		"Site",
		"Ramo de atividade",
		"Observações",
		"Apresentação",
		"Descrição de produtos/serviços",
		"Tratamento",
		"Orientação para atendimento",
		"Profissão",
		"Estado civil",
		"Data do cadastro",
		"ID da Unidade",
	}
	var rows [][]string
	for _, customer := range customers {
		rows = append(rows, []string{
			customer.CustomerId,
			customer.DueDate,
			customer.CompanyName,
			customer.TradeName,
			customer.Pronuntiation,
			customer.FundationDate,
			customer.CNPJ,
			customer.CPF,
			customer.RG,
			customer.IssuingAuthority,
			customer.FinanceEmail,
			customer.NotificationEmail,
			customer.Telephone,
			customer.Cellphone,
			customer.DDR,
			customer.Extension,
			customer.Mailbox,
			customer.Passport,
			customer.LateFee,
			customer.DayFee,
			customer.DelayDaysForBlocking,
			customer.ISSRetainer,
			customer.MunicipalInscription,
			customer.StateInscription,
			customer.Site,
			customer.BranchOfActivity,
			customer.Notes,
			customer.Presentation,
			customer.ServiceDescription,
			customer.Treatment,
			customer.ServiceOrientation,
			customer.Occuptation,
			customer.CivilState,
			customer.RegisterDate,
			customer.UnityID,
		})
	}
	t := table.New().
		Border(lipgloss.HiddenBorder()).
		Headers(columns...).
		Rows(rows...).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == 0 {
				return lipgloss.NewStyle().
					Foreground(lipgloss.Color("212")).
					Border(lipgloss.NormalBorder()).
					BorderTop(false).
					BorderLeft(false).
					BorderRight(false).
					BorderBottom(true).
					Bold(true)
			}
			if row%2 == 0 {
				return lipgloss.NewStyle().Foreground(lipgloss.Color("246"))
			}
			return lipgloss.NewStyle()
		})
	return t
}
