package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	quiz struct {
		time string `yaml:"time"`
	} `yaml:"quiz"`
}

type Config2 struct {
	time string `yaml:"time"`
}

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
