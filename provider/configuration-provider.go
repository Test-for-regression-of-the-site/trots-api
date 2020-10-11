package provider

import (
	"github.com/Test-for-regression-of-the-site/trots-api/constants"
	"github.com/spf13/viper"
	"log"
)

func LoadConfiguration() {
	viper.SetConfigName(constants.Trots)
	viper.SetConfigType(constants.Yml)
	viper.AddConfigPath(constants.Dot)
	if readingError := viper.ReadInConfig(); readingError != nil {
		log.Panicf("Reading configuration file error: %s \n", readingError)
	}
	log.Printf("Read configuration from file %s", viper.ConfigFileUsed())
}
