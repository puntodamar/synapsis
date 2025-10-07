package config

import (
	"log"
	"os"
)

type Config struct {
	HTTPAddress          string
	GRPCAddress          string
	GRPCInventoryAddress string
	DatabaseURL          string
	NATSUrl              string
}

func FromEnv() Config {
	c := Config{
		HTTPAddress:          get("HTTP_ADDRESS", ":8080"),
		GRPCAddress:          get("GRPC_ADDRESS", ":50051"),
		GRPCInventoryAddress: get("GRPC_INVENTORY_ADDRESS", "localhost:50051"),
		DatabaseURL:          get("DATABASE_URL", ""),
		NATSUrl:              get("NATS_URL", "nats://localhost:4222"),
	}
	if c.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
	return c
}

func get(k, def string) string {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	return v
}
