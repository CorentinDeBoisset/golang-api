package cmd

import (
	"fmt"
	"log"
	"net"
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

			network := viper.GetString("server.network")
			var listener net.Listener
			 if network == "tcp" {
				listener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port")))
				if err != nil {
					log.Printf("Could not register the TCP listener:\n\n    %s\n\n", err)
					os.Exit(1)
				}
			} else if network == "unix" {
				listener, err = net.Listen("unix", viper.GetString("server.socket_path"))
				if err != nil {
					log.Printf("Could not register the unix socket listener:\n\n    %s\n\n", err)
					os.Exit(1)
				}
			}

			http.Serve(listener, router)
		},
	})
}
