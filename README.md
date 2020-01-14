# Redis Info parser

Parses output from `INFO` redis command.

## Usage

```go
package main

import (
    "os"
    "io/ioutil"
    "json"
    "fmt"

    "github.com/geoffreybauduin/redis-info"
)

func main() {
    in, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        panic(err)
    }
    info, err := redisinfo.Parse(string(in))
    if err != nil {
        panic(err)
    }
    data, err := json.Marshal(info)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%s\n", data)
}
```
