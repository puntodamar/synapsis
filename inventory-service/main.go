package main

import (
	"fmt"
	"github.com/synapsis/common/config"
)

func main() {
	cfg := config.FromEnv()
	fmt.Printf("test air 22222")
	fmt.Printf("Starting inventory service at %s\n", cfg.HTTPAddress)
}
