package conf

import (
	"github.com/spf13/viper"
)

func init() {
	// Logger Defaults
	viper.SetDefault("logger.level", "info")
	viper.SetDefault("logger.encoding", "console")
	viper.SetDefault("logger.color", true)
	viper.SetDefault("logger.dev_mode", true)
	viper.SetDefault("logger.disable_caller", false)
	viper.SetDefault("logger.disable_stacktrace", true)

	// Pidfile
	viper.SetDefault("pidfile", "")

	// Profiler config
	viper.SetDefault("profiler.enabled", false)
	viper.SetDefault("profiler.host", "127.0.0.1")
	viper.SetDefault("profiler.port", "3001")

	// Server Configuration
	viper.SetDefault("server.host", "127.0.0.1")
	viper.SetDefault("server.port", "3000")
	viper.SetDefault("server.log_requests", true)
	viper.SetDefault("server.profiler_enabled", false)
	viper.SetDefault("server.profiler_path", "/debug")

	// Database Settings
	viper.SetDefault("storage.username", "postgres")
	viper.SetDefault("storage.password", "password")
	viper.SetDefault("storage.host", "localhost")
	viper.SetDefault("storage.port", "5432")
	viper.SetDefault("storage.database", "golang_api")
}
