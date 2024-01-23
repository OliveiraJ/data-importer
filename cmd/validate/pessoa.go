package validate

import (
	"fmt"

	"github.com/spf13/cobra"
)

func validatePessoa() *cobra.Command {
	pessoaCmd := &cobra.Command{Use: "pessoa",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("pessoa called")
		}}

	pessoaCmd.Flags().BoolVarP(&report, "relatório", "r", false, "Se um relatório de consistência deve ser gerado")

	return pessoaCmd
}
