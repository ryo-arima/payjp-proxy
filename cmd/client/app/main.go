package main

import (
	"os"

	"github.com/ryo-arima/payjp-proxy/pkg/client"
	"github.com/ryo-arima/payjp-proxy/pkg/config"
)

func main() {
	yamlConfig, err := config.LoadYaml(os.Getenv("PAYJP_PROXY_CONFIG_PATH"))
	if err != nil {
		panic(err)
	}
	conf, err := config.NewBaseConfig(yamlConfig)
	if err != nil {
		panic(err)
	}
	client.ClientForAppUser(*conf)
}
