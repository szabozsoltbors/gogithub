package config

import (
	"flag"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type env struct {
	PROFILE  string
	PORT     string
	LOGLEVEL string
}

// ENV - global environment settings
var ENV env

const configPath = "./environments"
const configType = "json"
const defaultProfile = "prod"

// InitEnv - initializes the environment profile
func InitEnv() error {
	profile := *readProfileFlag()

	viper.SetConfigName(profile)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)

	viper.SetDefault("PROFILE", "prod")
	viper.SetDefault("PORT", ":8080")
	viper.SetDefault("LOGLEVEL", "info")

	err := viper.ReadInConfig()

	if err != nil {
		log.WithError(err).Error("Error reading config file")
		log.Warning("Fallback to default settings")
		return nil
	}

	err = viper.Unmarshal(&ENV)
	if err != nil {
		log.WithError(err).Error("Unable to decode the config file into struct")
		return err
	}

	return nil
}

func readProfileFlag() *string {
	profile := flag.String("profile", "prod", "Choose a profile: prod or dev")
	flag.Parse()

	if *profile == "" {
		*profile = defaultProfile
	}

	return profile
}
