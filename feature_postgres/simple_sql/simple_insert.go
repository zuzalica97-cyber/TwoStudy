package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UsersInsertRow(
	conn *pgx.Conn,
	ctx context.Context,
	u UserModel,
) error {

	sqlQuery := `
	INSERT INTO users (name, money) 
	VALUES ($1, $2);
	`

	_, err := conn.Exec(ctx, sqlQuery, u.Name, u.Money)

	return err
}

func ProductsInsertRow(
	conn *pgx.Conn,
	ctx context.Context,
	p ProductModel,
) error {
	sqlQuery := `
	INSERT INTO products (name, description, cost, amount) 
	VALUES ($1, $2, $3, $4);
	`
	_, err := conn.Exec(ctx, sqlQuery, p.Name, p.Description, p.Cost, p.Amount)

	return err
}
