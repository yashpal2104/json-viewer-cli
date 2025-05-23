/*
Copyright © 2025 Yash  <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yashpal2104/json-viewer-cli/main"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "It helps you select fruits of different sizes and test out the Renovate tool",
	Long: `This CLI is designed for you to select fruits based on different sizes blah blah blah you don't care about this
	 	even I don't care about it what we are trying to do here is to test out
			the Renovate tool to automate dependency updates and patch regular security threats through making PRs to fix it.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { 

	// },
}

var ListCmd = &cobra.Command{
	Use: "list",
	Short: "lists all the fruits",
	Run: func (cmd *cobra.Command, args[] string)  {
		for i, fruit = range main.FileParsedData{
			
		} 
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.json-viewer-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand()
}


