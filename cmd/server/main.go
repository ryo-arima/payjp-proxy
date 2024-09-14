package main

import (
	"os"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/server"
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
	server.Main(*conf)
}
