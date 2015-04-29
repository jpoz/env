## env_struct

Decode Structs from the Environment

```golang
package main

import (
  "github.com/jpoz/env_struct"
)

type Config struct {
  Host string `env:"HOST" default"localhost"`
  Env string `env:"ENV" default"development"`
}

func main() {
  config := Config{}

  env_struct.Decode(config)

  fmt.Printf("%s running in %s", config.Host, config.Env)
}
```

```
$ ENV=dev go run main.go
localhost running in dev
```
