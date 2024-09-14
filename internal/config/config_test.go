package main

import (
	"github.com/spf13/viper"
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestConfigLoad(t *testing.T) {
	viper.SetConfigFile("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		t.Fatalf("Failed to read config file: %v", err)
	}

	err = viper.Unmarshal(&Config{})
	if err != nil {
		t.Fatalf("Failed to unmarshal config: %v", err)
	}

	validate := validator.New()
	err = validate.Struct(config)
	if err != nil {
		t.Fatalf("Failed to validate config: %v", err)
	}
}