package validate

import (
	"fmt"

	"github.com/spf13/cobra"
)

// clienteCmd represents the cliente command
// var clienteCmd = &cobra.Command{
// 	Use:   "cliente",
// 	Short: "A brief description of your command",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("cliente called")
// 		fmt.Println(report)

// 	},
// }

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
			fmt.Println("cliente called")
			fmt.Println(report)

		},
	}

	clienteCmd.Flags().BoolVarP(&report, "relatório", "r", false, "Se um relatório de consistência deve ser gerado")

	return clienteCmd
}

// func init() {
// 	validateCmd.AddCommand(clienteCmd)
// 	clienteCmd.Flags().BoolVarP(&report, "relatório", "r", false, "Se um relatório de consistência deve ser gerado")
// 	// Here you will define your flags and configuration settings.

// 	// Cobra supports Persistent Flags which will work for this command
// 	// and all subcommands, e.g.:
// 	// clienteCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	// Cobra supports local flags which will only run when this command
// 	// is called directly, e.g.:
// 	// clienteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
