package cmd

import (
	"log"
	"os"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Kudo CMS - API Gateway",
	Short: "Kudo CMS - API Gateway",
	Long:  `Kudo CMS - API Gateway`,
}

// Execute the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(0)
	}
}

func init() {
	cobra.OnInitialize(func() {
		viper.SetConfigFile("./config/default.json")
		if err := viper.ReadInConfig(); err != nil {
			log.Println(err)
			os.Exit(0)
		}
	})
}
