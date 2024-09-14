package main

import (
	"fmt"
	"log"
	"os"
	"tx_parser/config"
)

func main() {
	logger := log.New(os.Stdout, "CONFIG-LOADER: ", log.Ldate|log.Ltime|log.Lshortfile)

	c, err := config.LoadConfig("config/config.json", logger)
	if err != nil {
		log.Panicf("cannot read config file")
	}

	fmt.Println(c.Server.Host)
	fmt.Println(c.Server.Port)
	fmt.Println(c.Client.Endpoint)
}
