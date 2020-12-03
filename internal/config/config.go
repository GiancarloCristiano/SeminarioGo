package config

import ( 
	"io/ioutil"
	"gopkg.in/yaml.v3"
)

type DbConfig struct {
	Type string `yaml:"type"`
	Driver string `yaml:"driver"`
	Conn string `yaml:"conn"`
}

type Config struct {
	DB DbConfig `yaml:"db"`
	Version string `yaml:"version"`
}

func LoadConfig(filename string) (*Config, error) {
	
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var c = &Config{}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}