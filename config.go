package main

import (
	"log"

	"github.com/spf13/viper"
)

type Configurations struct{
	AuthToken string
}

var C Configurations

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