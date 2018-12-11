package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/globalsign/mgo"
	"github.com/spf13/viper"
)

type (
	// Constants is set according to the environment variables given in config.toml
	// config.toml must conform to the structure of this. PORT will default to
	// 8080 if it's not set in config.toml.
	Constants struct {
		PORT  string
		Mongo struct {
			URL    string
			DBName string
		}
	}

	// Config represents the actual configuration object passed around. It includes
	// the constants set in Constants and an actual MongoDB instance.
	Config struct {
		Constants
		Database *mgo.Database
	}
)

// New is used to generate a configuration instance which will be passed around
// the codebase.
func New(configFilename string) (*Config, error) {
	config := Config{}
	constants, err := initConstants(configFilename)
	config.Constants = constants
	if err != nil {
		return &config, err
	}

	dbSession, err := mgo.Dial(config.Constants.Mongo.URL)
	if err != nil {
		return &config, err
	}

	config.Database = dbSession.DB(config.Constants.Mongo.DBName)
	return &config, err
}

func initConstants(configFilename string) (Constants, error) {
	// Config filename without the .TOML or .YAML extension
	viper.SetConfigName(configFilename)
	// Search server project directory (NOT this directory)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return Constants{}, err
	}

	// Watch for changes to the configuration file and recompile
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	// Set default PORT to 8080 (if not given) and read the config
	viper.SetDefault("PORT", "8080")
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Error reading config file, %s", err)
	}

	var constants Constants
	return constants, viper.Unmarshal(&constants)
}
