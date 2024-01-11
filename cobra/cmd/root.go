package cmd

import (
	"cobra/helper"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var reverseCmd = &cobra.Command{
	Use:     "reverse",
	Short:   "Reverses a string",
	Aliases: []string{"rev"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := helper.Reverse(args[0])
		fmt.Println(res)
	},
}

var uppercaseCmd = &cobra.Command{
	Use:     "uppercase",
	Short:   "Uppercase a string",
	Aliases: []string{"upper"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := helper.Uppercase(args[0])
		fmt.Println(res)
	},
}

var option bool
var modifyCmd = &cobra.Command{
	Use:     "modify",
	Short:   "Modify a string",
	Aliases: []string{"modif"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := helper.Modify(args[0], option)
		fmt.Println(res)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	modifyCmd.Flags().BoolVarP(&option, "option", "o", false, "Modify option")
	rootCmd.AddCommand(reverseCmd)
	rootCmd.AddCommand(uppercaseCmd)
	rootCmd.AddCommand(modifyCmd)
}
