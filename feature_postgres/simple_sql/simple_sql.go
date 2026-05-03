package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateProductTable(conn *pgx.Conn, ctx context.Context) error { //создаем таблицу products, если она не существует
	sqlQwery := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL, 
		description TEXT NOT NULL,
		cost INTEGER NOT NULL,
		amount INTEGER NOT NULL,
		productid INTEGER NOT NULL
	);
	`

	_, err := conn.Exec(ctx, sqlQwery)

	return err
}
