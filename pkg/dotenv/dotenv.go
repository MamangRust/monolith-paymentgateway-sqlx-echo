package dotenv

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Viper reads environment variables from a configuration file based on the
// value of the APP_ENV environment variable. The configuration file is read
// using the viper package. If the APP_ENV variable is not set, the default
// value is "development".
//
// The configuration file is read from the following locations:
// - development: .env
// - docker: /app/docker.env
// - production: /app/production.env
// - kubernetes: no configuration file is read
// - test: no configuration file is read
//
// If an error occurs while reading the configuration file, an error is
// returned.
func Viper() error {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	var configFile string
	useConfigFile := true

	switch env {
	case "docker":
		configFile = "/app/docker.env"
	case "production":
		configFile = "/app/production.env"
	case "kubernetes", "test":
		useConfigFile = false
	default:
		configFile = ".env"
	}

	viper.AutomaticEnv()

	if useConfigFile {
		viper.SetConfigFile(configFile)

		err := viper.ReadInConfig()
		if err != nil {
			return fmt.Errorf("error reading config file %s: %w", configFile, err)
		}
	}

	return nil
}
