package main

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
)

type Config map[string]string

func getConfig() Config {
	configFile := flag.String("env", ".env", "environment file path")
	if isTesting {
		configFile = flag.String("env", ".env.test", "environment file path")
	}
	flag.Parse()
	log.Println("CONFIG", *configFile)
	config, err := godotenv.Read(*configFile)
	printErrIfNotNil(err)
	return config
}
