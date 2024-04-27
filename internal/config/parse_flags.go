package config

import "flag"

func ParseFlags() {
	flag.IntVar(&Config.Port, "port", 3000, "Port for the server to listen on")
	flag.StringVar(&Config.Env, "env", "development", "Environment the server is being run in, e.g. 'development'")
	flag.Parse()
}
