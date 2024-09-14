package config

type PayJpConfig struct {
	Payjp string `yaml:"Payjp"`
}

type Payjp struct {
	Secret string `yaml:"secret"`
}
