package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/corentindeboisset/golang-api/app/router"
)

// Version command
func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "api",
		Short: "Start the API",
		Long:  `Start the API`,
		Run: func(cmd *cobra.Command, args []string) {
			router, err := router.GetRouter()
			if err != nil {
				log.Printf("The launch of the API crashed with the following error:\n\n    %s\n\n", err)
				os.Exit(1)
			}

			http.ListenAndServe(
				fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port")),
				router,
			)
		},
	})
}
