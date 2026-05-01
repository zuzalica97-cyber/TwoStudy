package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UserDeleteRow(conn *pgx.Conn, ctx context.Context, task []int) error { //удаляем строки, id которых совпадает с id из слайса task

	sqlQuery := `
	DELETE FROM users
	WHERE id = ANY($1);
	`

	_, err := conn.Exec(ctx, sqlQuery, task)

	return err
}

func ProductDeleteRow(conn *pgx.Conn, ctx context.Context, task []int) error { //удаляем строки, id которых совпадает с id из слайса task

	sqlQuery := `
	DELETE FROM products
	WHERE id = ANY($1);
	`

	_, err := conn.Exec(ctx, sqlQuery, task)

	return err
}
