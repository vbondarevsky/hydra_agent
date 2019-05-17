package main

import (
	"./api"
	"./config"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
)

func initConfig() (config.Configuration, error) {
	pflag.String("config", ".", "path to config file")
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)

	viper.SetConfigName("config")
	viper.AddConfigPath(viper.GetString("config"))

	var configuration config.Configuration

	err := viper.ReadInConfig()
	err = viper.Unmarshal(&configuration)

	return configuration, err
}

func main() {

	settings, err := initConfig()
	if err != nil {
		panic(err)
	}

	server := fmt.Sprintf(":%d", settings.Api.Port)
	http.HandleFunc("/health", api.HealthHandler)
	http.HandleFunc("/platform", api.PlatformHandler)

	if err := http.ListenAndServe(server, nil); err != nil {
		panic(err)
	}
}
