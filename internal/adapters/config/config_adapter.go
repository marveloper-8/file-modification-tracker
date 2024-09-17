package config

import (
	"github.com/spf13/viper"
	"log"
)

type ConfigAdapter struct {
	directory string
	checkFreq int
}

func NewConfigAdapter() *ConfigAdapter {
	return &ConfigAdapter{}
}

func (c *ConfigAdapter) LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
	c.directory = viper.GetString("directory")
	c.checkFreq = viper.GetInt("check_freq")
}

func (c *ConfigAdapter) GetDirectory() string {
	return Config.Directory
}

func (c *ConfigAdapter) GetCheckFrequency() int {
	return Config.CheckFreq
}
