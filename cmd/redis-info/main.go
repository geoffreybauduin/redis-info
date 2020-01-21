package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/geoffreybauduin/redis-info"
)

func main() {
	v, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	info, err := redisinfo.Parse(string(v))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", info)
}
