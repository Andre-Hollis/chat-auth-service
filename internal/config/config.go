package config

import (
	"encoding/json"
	"os"

	"github.com/Andre-Hollis/chat-auth-service/internal/models"
)

func LoadConfig(file string) (models.Config, error) {
	var config models.Config

	data, err := os.ReadFile(file)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
