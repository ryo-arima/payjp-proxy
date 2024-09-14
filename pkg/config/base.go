package config

type BaseConfig struct {
	YamlConfig YamlConfig
}

func NewBaseConfig(conf YamlConfig) (*BaseConfig, error) {
	return &BaseConfig{
		YamlConfig: conf,
	}, nil
}
