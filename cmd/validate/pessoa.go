/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package validate

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pessoaCmd represents the pessoa command
// var pessoaCmd = &cobra.Command{
// 	Use:   "pessoa",
// 	Short: "A brief description of your command",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("pessoa called")
// 	},
// }

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

// func init() {
// 	validateCmd.AddCommand(pessoaCmd)
// 	pessoaCmd.Flags().BoolVarP(&report, "relatório", "r", false, "Se um relatório de consistência deve ser gerado")
// 	// Here you will define your flags and configuration settings.

// 	// Cobra supports Persistent Flags which will work for this command
// 	// and all subcommands, e.g.:
// 	// pessoaCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	// Cobra supports local flags which will only run when this command
// 	// is called directly, e.g.:
// 	// pessoaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
