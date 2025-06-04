package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("resources")
	fmt.Println(viper.Get("quiz.time"))
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
