package main

import (
	"fmt"

	"github.com/jpoz/env/decoder"
)

type Config struct {
	Env     string `env:"ENV"`
	Addr    string `expand:"$HOST:$PORT"`
	Workers int    `env:"WORKERS"`
}

func main() {
	config := Config{}

	decoder.Decode(&config)

	fmt.Printf("%s running in %s with %d workers",
		config.Addr, config.Env, config.Workers)
}
