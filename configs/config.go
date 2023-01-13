package configs

type Config struct {
	Client struct {
		Token    string `yaml:"token"`
		Endpoint struct {
			URL  string `yaml:"url"`
			Port int    `yaml:"port"`
		} `yaml:"endpoint"`
	} `yaml:"client"`
}
