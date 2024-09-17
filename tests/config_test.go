package tests

import (
	"file-modification-tracker/internal/adapters/config"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfigAdapter_LoadConfig(t *testing.T) {
	viper.Set("directory", "./test-dir")
	viper.Set("check_freq", 60)

	adapter := config.NewConfigAdapter()
	adapter.LoadConfig()

	assert.Equal(t, "./test-dir", adapter.GetDirectory())
	assert.Equal(t, 60, adapter.GetCheckFrequency())
}
