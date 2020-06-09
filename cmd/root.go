package cmd

import (
	"log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	// "go.uber.org/zap"
	"github.com/corentindeboisset/golang-api/conf"
)

var (
	// Config file
	cfgFile		string

	rootCmd = &cobra.Command{
		Use:   conf.Executable,
		Short: "My Recipes backend",
		Long: `My Recipes backend, a website to manage lists of recipes and ingredients`,
	}
)

// Execute the rootCommand
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: ./config.yml, /etc/config.yml)")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search current directory
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/")
		viper.SetConfigName("config")
	}

	viper.SetEnvPrefix("golang_api")
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		// TODO replace log by zap
		log.Fatal(err)
	}
	// TODO debug-log the full path to the used config file
}
