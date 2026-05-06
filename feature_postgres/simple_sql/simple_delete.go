package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UserDeleteRow(conn *pgx.Conn, ctx context.Context, task []int) error { //удаляем строки, id которых совпадает с id из слайса task

	sqlQuery := `
	DELETE FROM users
	WHERE userid = ANY($1);
	`

	_, err := conn.Exec(ctx, sqlQuery, task)

	return err
}

func ProductDeleteRow(conn *pgx.Conn, ctx context.Context, task []int) error { //удаляем строки, id которых совпадает с id из слайса task

	sqlQuery := `
	DELETE FROM products
	WHERE productidp = ANY($1);
	`

	_, err := conn.Exec(ctx, sqlQuery, task)

	return err
}

func BaseDeleteRow(conn *pgx.Conn, ctx context.Context, task []int) error { //удаляем строки, id которых совпадает с id из слайса task

	sqlQuery := `
	DELETE FROM bases
	WHERE baseid = ANY($1);
	`

	_, err := conn.Exec(ctx, sqlQuery, task)

	return err

}
