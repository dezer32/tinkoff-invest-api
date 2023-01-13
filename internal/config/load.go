package config

import (
	"github.com/dezer32/tinkoff-invest-api/configs"
	"github.com/go-yaml/yaml"
	"os"
)

func Load(files ...string) (config *configs.Config, err error) {
	config = new(configs.Config)

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return config, err
		}
		defer f.Close()

		decoder := yaml.NewDecoder(f)
		err = decoder.Decode(config)
		if err != nil {
			return config, err
		}
	}

	return
}
