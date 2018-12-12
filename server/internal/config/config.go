package config

import (
	"github.com/bli940505/threaded/server/internal/types"
	"github.com/spf13/viper"
)

// NewConfig TODO
func NewConfig(filename string) (*types.Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(filename)
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 4000)
	viper.ReadInConfig()

	var env types.Config
	return &env, viper.Unmarshal(&env)
}
