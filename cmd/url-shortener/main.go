package main

import (
	"fmt"
	"link-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	// TODO: init logger

	// TODO: init storage

	// TODO: init router

	// TODO: init server
}
