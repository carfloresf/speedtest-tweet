package config

import (
	"github.com/spf13/viper"
	"log"
)

type Twitter struct {
	AccessToken    *string `mapstructure:"accessToken"`
	AccessSecret   *string `mapstructure:"accessSecret"`
	ConsumerKey    *string `mapstructure:"consumerKey"`
	ConsumerSecret *string `mapstructure:"consumerSecret"`
}

type Config struct {
	Twitter          `mapstructure:"twitterConfig"`
	Threshold        *float64 `mapstructure:"threshold"`
	ExpectedDownload *int     `mapstructure:"expectedDownload"`
	AtTwitter        *string  `mapstructure:"atTwitter"`
}

func GetConfiguration() *Config {
	configuration := Config{}

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	if *configuration.ConsumerKey == "" || *configuration.ConsumerSecret == "" || *configuration.AccessToken == "" || *configuration.AccessSecret == "" {
		log.Fatalln("Consumer key/secret and Access token/secret required")
	}

	return &configuration
}
