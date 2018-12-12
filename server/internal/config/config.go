package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config defaults to localhost:4000
type Config struct {
	Host string
	Port string
}

// NewConfig TODO
func NewConfig(filename string) (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(filename)
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 4000)
	viper.ReadInConfig()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed at ", e.Name)
	})

	var env Config
	return &env, viper.Unmarshal(&env)
}

// PrintInfo prints the config info of this config
func (c *Config) PrintInfo() {
	fmt.Println("---Config Information---")
	log.Printf("Host: %s\n", c.Host)
	log.Printf("Port: %s\n", c.Port)
}
