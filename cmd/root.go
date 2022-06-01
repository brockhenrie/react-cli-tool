/*
Copyright © 2022 Brock Henrie bhenriedev@gmail.com
*/
package cmd

import (
	"os"
	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "react",
	Short: "A Component CLI Tool",
	Long: "A CLI tool that builds react app components with either .jsx or .tsx",	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.joke.yaml)")

//    rootCmd.Flags.BoolP("typescript","ts", false,"Use flag to change output type to typescript")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
//	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



