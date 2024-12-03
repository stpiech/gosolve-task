package main

import (
	"github.com/stpiech/gosolve-task/internal/api"
	"github.com/stpiech/gosolve-task/internal/loader"
	"github.com/stpiech/gosolve-task/internal/logger"
)

func main() {
	config, err := loader.LoadConfig("config.json")

	if err != nil {
		panic(err)
	}

	values, err := loader.FileToSlice("input.txt")

	if err != nil {
		panic(err)
	}

	err = logger.SetLogLevel(config.LogLevel)

	if err != nil {
		panic(err)
	}

	err = api.RegisterSearchValueEndpoint(values, config.Port)

	if err != nil {
		panic(err)
	}
}
