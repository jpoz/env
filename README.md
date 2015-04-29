## env_struct

Decode Structs from the Environment

```go
package main

import (
  "github.com/jpoz/env_struct"
)

type Config struct {
  Env string `env:"ENV"`
  Addr string `expand:"$HOST:$PORT"`
}

func main() {
  config := Config{}

  env_struct.Decode(config)

  fmt.Printf("%s running in %s", config.Addr, config.Env)
}
```

```
$ PORT=9999 ENV=dev go run main.go
:9999 running in dev
```
