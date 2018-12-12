package utils

import (
	"github.com/bli940505/threaded/server/internal/types"
	"github.com/go-pg/pg"
	"github.com/spf13/viper"
)

// ReadConfig reads both database.toml and env.toml under internal/configs/
// both config file can be specifying by passing in an argument flag to main()
// both have defaults but only sets the Database field will be set in
// pg.Options
func ReadConfig(envFilename, databaseFilename string) (*types.URL, *pg.Options, error) {
	viper.AddConfigPath("./internal/configs")

	// read env
	viper.SetConfigName(envFilename)
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 4000)
	viper.ReadInConfig()

	var env types.URL
	if err := viper.Unmarshal(&env); err != nil {
		return &env, &pg.Options{Database: "threaded"}, err
	}

	// read database
	viper.SetConfigName(databaseFilename)
	viper.SetDefault("database", "threaded")
	viper.ReadInConfig()

	var opt pg.Options
	return &env, &opt, viper.Unmarshal(&opt)
}
