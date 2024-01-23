package validate

import (
	"fmt"

	"github.com/spf13/cobra"
)

func validateFornecedor() *cobra.Command {
	fornecedorCmd := &cobra.Command{
		Use:   "fornecedor",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("fornecedor called")
		},
	}

	fornecedorCmd.Flags().BoolVarP(&report, "relatório", "r", false, "Se um relatório de consistência deve ser gerado")

	return fornecedorCmd
}
