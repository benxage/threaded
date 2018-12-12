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

// New TODO
func New(filename string) *Config {
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
	viper.Unmarshal(&env)
	return &env
}

// PrintInfo prints the config info of this config
func (c *Config) PrintInfo() {
	fmt.Println("---Config Information---")
	log.Printf("Host: %s\n", c.Host)
	log.Printf("Port: %s\n", c.Port)
}
