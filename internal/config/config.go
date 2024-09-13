package main

import (
    "github.com/spf13/viper"
    "github.com/go-playground/validator/v10"
    "log"
)

type Config struct {
	Directory string `validate:"required,dir"`
	Frequency int `validate:"required,gt=0"`
	APIEndpoint string `validate:"required,url"`
}

var config Config

func NewConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}

	validate := validator.New()
	err = validate.Struct(config)
	if err != nil {
		log.Fatalf("Error validating config: %v", err)
	}

	return &config
}
