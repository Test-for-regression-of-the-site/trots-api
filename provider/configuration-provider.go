package provider

import (
	"github.com/Test-for-regression-of-the-site/trots-api/configuration"
	"github.com/Test-for-regression-of-the-site/trots-api/constants"
	"github.com/spf13/viper"
	"log"
)

var Configuration = LoadConfiguration()

func LoadConfiguration() configuration.Configuration {
	viper.SetConfigName(constants.Trots)
	viper.SetConfigType(constants.Yml)
	viper.AddConfigPath(constants.Dot)
	if readingError := viper.ReadInConfig(); readingError != nil {
		log.Panicf("Reading configuration file error: %s \n", readingError)
	}
	log.Printf("Read configuration from file %s", viper.ConfigFileUsed())
	return configuration.Configuration{
		Address: viper.GetString(constants.ServerAddressKey),
		Timeout: viper.GetDuration(constants.TimeoutKey),
		Lighthouse: configuration.LighthouseConfiguration{
			Image:             viper.GetString(constants.LighthouseImageKey),
			Tag:               viper.GetString(constants.LighthouseTagKey),
			ReportsTargetPath: viper.GetString(constants.LighthouseReportsTargetPathKey),
			ReportsSourcePath: viper.GetString(constants.LighthouseReportsSourcePathKey),
		},
		Mongo: configuration.MongoConfiguration{
			Address: viper.GetString(constants.MongoAddressKey),
			Timeout: viper.GetDuration(constants.MongoTimeoutKey),
		},
	}
}
