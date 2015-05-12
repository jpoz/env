## env/decoder
Decode Structs from the Environment

[![GoDoc](https://godoc.org/github.com/jpoz/env/decoder?status.svg)](https://godoc.org/github.com/jpoz/env/decoder)

```go
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
```

```
$ HOST=hello PORT=9000 ENV=dev WORKERS=123 go run example.go
hello:9000 running in dev with 123 workers
```
