package tests

import (
	"file-modification-tracker/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config.LoadConfig()

	assert.NotEmpty(t, config.Config.Directory, "Directory should not be empty")
	assert.NotEmpty(t, config.Config.CheckFreq, "Check frequency should not be empty")
	assert.NotEmpty(t, config.Config.RemoteAPI, 0, "Remote API should not be empty")
}

func TestInvalidConfig(t *testing.T) {
	config.Config.CheckFreq = -1
	validate := config.Config.ValidateConfig()
	assert.NotNil(t, validate, "CheckFreq validation should fail for negative value")
}