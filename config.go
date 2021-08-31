package main

import (
	"log"

	"github.com/spf13/viper"
)

type Configurations struct{
	AuthToken string
	Server    ServerConfigurations
	Database  DatabaseConfigurations
}

type ServerConfigurations struct {
	Port int
}
type DatabaseConfigurations struct {
	DbName     string
	DbUser     string
	DbPassword string
	DbPort     int
	DbBaseUrl  string
}

var C Configurations
// loading configurations from conf.yaml file using viper
func loadConfigurations(){
	viper.AddConfigPath(".")
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}