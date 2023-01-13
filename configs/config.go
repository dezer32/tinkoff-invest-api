package configs

import "reflect"

type Config struct {
	Client struct {
		Token    string `yaml:"token"`
		Endpoint struct {
			URL  string `yaml:"url"`
			Port int    `yaml:"port"`
		} `yaml:"endpoint"`
	} `yaml:"client"`
	Mapper map[string]reflect.Type `yaml:"mapper"`
}
