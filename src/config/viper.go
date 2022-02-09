package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("./")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: \n %w ", err))
	}

}

func GETSTRING(key string) string {
	return viper.GetString(key)
}
