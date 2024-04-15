package config

import "flag"

type AppConfig struct {
	Port int
	Env  string
}

func ParseConfig() *AppConfig {
	var cfg AppConfig
	flag.IntVar(&cfg.Port, "port", 3000, "Port for the server to listen on")
	flag.StringVar(&cfg.Env, "env", "development", "Environment the server is being run in, e.g. 'development'")
	flag.Parse()

	return &cfg
}
