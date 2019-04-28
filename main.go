package main

import (
	"./config"
	"encoding/json"
	"fmt"
	"github.com/matishsiao/goInfo"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"

	"github.com/inconshreveable/go-update"
)

type System struct {
	Kernel   string `json:"kernel"`
	Core     string `json:"core"`
	Platform string `json:"platform"`
	OS       string `json:"os"`
	Hostname string `json:"hostname"`
	CPUs     int    `json:"cpus"`
}

type Health struct {
	Version string `json:"version"`
	System  `json:"system"`
}

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	gi := goInfo.GetInfo()
	health := Health{"alpha", System{gi.Kernel, gi.Core, gi.Platform, gi.OS, gi.Hostname, gi.CPUs}}
	buf, _ := json.Marshal(health)

	fmt.Fprintf(w, "%s", buf)
}

func handlerAutoUpdate(w http.ResponseWriter, r *http.Request) {
	newVersion, err := os.Open(`D:\main.exe`)
	if err != nil {
		// error handling
		log.Fatalf(err.Error())
	}
	defer newVersion.Close()
	err = update.Apply(newVersion, update.Options{})
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Fprintf(w, "update success")
}

func initConfig() (config.Configuration, error) {
	pflag.String("config", "config.yml", "config file")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	configFile := viper.GetString("config")
	log.Printf("flag is %s", configFile)

	viper.SetConfigName("config")
	viper.AddConfigPath(configFile)
	viper.AddConfigPath(".")

	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	log.Printf("database uri is %s", configuration.Database.ConnectionUri)
	log.Printf("port for this application is %d", configuration.Server.Port)

	return configuration, err
}

func main() {

	settings, err := initConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	server := fmt.Sprintf(":%d", settings.Server.Port)
	http.HandleFunc("/health", handlerHealth)
	http.HandleFunc("/autoupdate", handlerAutoUpdate)
	err = http.ListenAndServe(server, nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
