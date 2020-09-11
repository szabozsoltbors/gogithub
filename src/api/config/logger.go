package config

import (
	log "github.com/sirupsen/logrus"
)

// InitLogger - initializes the logrus logger
func InitLogger() error {

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	logLevel, err := log.ParseLevel(ENV.LOGLEVEL)

	if err != nil {
		log.WithError(err).Error("Unable to parse loglevel")
		return err
	}

	log.SetLevel(logLevel)
	return nil
}
