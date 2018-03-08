package cmd

import (
	"github.com/fahribaharudin/api_gateway/app"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve, start the server app to serve http",
	Long:  `Serve, start the server app to serve http`,
	Run: func(cmd *cobra.Command, args []string) {
		var app = app.Kernel{}
		// app.Init() // wrapping up some components together
		// app.ParseSwaggerAPIEndpoints()
		// app.RegisterRoutes()

		app.Bootstrap()
		app.Run() // waking up the monster!
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
