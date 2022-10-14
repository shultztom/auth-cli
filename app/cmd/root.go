package cmd

import (
	"auth-cli/app/utils"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var appName string

func init() {
	rootCmd.PersistentFlags().StringVar(&appName, "appName", appName, "App Name you want token for")
	viper.BindPFlag("appName", rootCmd.PersistentFlags().Lookup("appName"))
}

var rootCmd = &cobra.Command{
	Use:   "auth-cli",
	Short: "CLI for working with auth-api-go",
	Long:  `CLI for working with auth-api-go - https://github.com/shultztom/auth-api-go`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(appName) == 0 {
			fmt.Println("Please add --appName Parameter")
			return
		}

		token := utils.CreateToken(appName)
		fmt.Println(token)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
