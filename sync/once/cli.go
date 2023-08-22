package main

import "sync"

var once sync.Once

var config *Config

type Config struct {
}

func initialize() {
	config = &Config{}
}

func GetConfig() *Config {
	once.Do(initialize)
	return config
}
