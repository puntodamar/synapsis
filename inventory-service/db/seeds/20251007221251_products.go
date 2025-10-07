package seeds

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"

	"github.com/pressly/goose/v3"
)

func init() { goose.AddMigrationContext(upSeedProducts, downSeedProducts) }

//go:embed fixtures/products.json
var seedProductsJSON []byte

type Product struct {
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Stock int32  `json:"stock"`
}

func upSeedProducts(ctx context.Context, tx *sql.Tx) error {
	var items []Product
	if err := json.Unmarshal(seedProductsJSON, &items); err != nil {
		return err
	}
	for _, p := range items {
		if _, err := tx.ExecContext(ctx, `
			INSERT INTO products (sku, name, stock)
			VALUES ($1, $2, $3)
			ON CONFLICT (sku) DO UPDATE
			SET name = EXCLUDED.name,
			    stock = EXCLUDED.stock
		`, p.SKU, p.Name, p.Stock); err != nil {
			return err
		}
	}
	return nil
}

func downSeedProducts(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `DELETE FROM products WHERE sku IN (SELECT sku FROM products)`)
	return err
}
