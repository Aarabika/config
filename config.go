package config

import (
	"github.com/spf13/viper"
	"strings"
)

func Run() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func GetConfig(name string) string {
	return viper.GetString(name)
}
