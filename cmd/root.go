package cmd

import (
	"balance/helper"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "balance",
	Short: "A brief description of your application",
	Long:  `A longer description`,
}

var balanceCmd = &cobra.Command{
	Use:     "balance",
	Short:   "Get balance by account address",
	Aliases: []string{"b"},
	Args:    cobra.ExactArgs(1),
	// Run: func(cmd *cobra.Command, args []string) {
	// 	cmd.Run(cmd, args)
	// },
	Run: helper.Run(),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(balanceCmd)
}
