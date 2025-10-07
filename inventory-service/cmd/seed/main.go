// services/inventory-service/cmd/migrate/main.go
package main

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/synapsis/common/config"
	_ "github.com/synapsis/inventory-service/db/seeds"
	"log"
)

func main() {
	cfg := config.FromEnv()
	db, err := sql.Open("pgx", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	if err := goose.UpContext(ctx, db, "db/seeds"); err != nil {
		log.Fatal("seeding failed:", err)
	}
	log.Println("seeding applied successfully")
}
