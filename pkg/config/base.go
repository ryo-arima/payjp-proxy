package config

type BaseConfig struct {
	YamlConfig  YamlConfig
	PayJpConfig PayjpConfig
}

func NewBaseConfig(conf YamlConfig) (*BaseConfig, error) {
	payjpConfig, err := NewPayjpConfig(conf)
	if err != nil {
		return nil, err
	}

	return &BaseConfig{
		YamlConfig:  conf,
		PayJpConfig: payjpConfig,
	}, nil
}
