package config

import (
	"github.com/payjp/payjp-go/v1"
)

type PayjpConfig struct {
	Service *payjp.Service
}

func NewPayjpConfig(conf YamlConfig) (*PayjpConfig, error) {
	secret := conf.Payjp.Secret
	service := payjp.New(secret, nil)
	return &PayjpConfig{
		Service: service,
	}, nil
}
