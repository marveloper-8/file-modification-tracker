package config

import (
    "github.com/spf13/viper"
    "log"
)

type Config struct {
	DirectoryToMonitor string
	CheckFrequency int
	APIEndpoint string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resources")
	
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
		return nil, err
	}

	return &config, nil
}