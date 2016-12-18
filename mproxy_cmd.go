package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/shanmuganandh/go-mysql-proxy/mproxy"
	"os"
)

var conf mproxy.Config

func main() {
	if _, confErr := toml.DecodeFile("sample-config.toml", &conf); confErr != nil {
		fmt.Println("Error in parsing the config file")
		os.Exit(1)
	}

	if router, err := mproxy.NewRouter(conf); err == nil {
		router.Start()
	} else {
		fmt.Println(err)
	}
}
