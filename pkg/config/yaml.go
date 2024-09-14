package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

type YamlConfig struct {
	Application Application `yaml:"Application"`
}

type Server struct {
	Port  string `yaml:"port"`
	Admin Admin  `yaml:"admin"`
}

type Client struct {
	ServerEndpoint string `yaml:"ServerEndpoint"`
	UserEmail      string `yaml:"UserEmail"`
	UserPassword   string `yaml:"UserPassword"`
	HomeDir        string `yaml:"HomeDir"`
}

type Application struct {
	Server Server `yaml:"Server"`
	Client Client `yaml:"Client"`
}

type Admin struct {
	Emails []string `yaml:"emails"`
}

type YamlLoader interface {
	LoadYaml(filePath string) (YamlConfig, error)
}

func LoadYaml(filePath string) (YamlConfig, error) {
	var config YamlConfig
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
