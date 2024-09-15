package config

	import (
	"log"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var Config Settings

type Settings struct {
	Directory string `validate:"required,dir"`
	CheckFreq int `validate:"required,min=1"`
	RemoteAPI string `validate:"required,url"`
	TimerInterval int `mapstructure:"remote_api"`
}

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	validate := validator.New()
	err = validate.Struct(Config)
	if err != nil {
		log.Fatalf("Failed to validate config: %v", err)
	}
}