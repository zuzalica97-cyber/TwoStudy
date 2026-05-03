package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UserMoneyUpdateRow(conn *pgx.Conn, ctx context.Context, moneyU int, idU int) error { //обновляем строку, id которой совпадает с idU, устанавливая значение money равным moneyU

	sqlQuery := `
	UPDATE users
	SET usermoney = $1
	WHERE userid = $2;
	`

	_, err := conn.Exec(ctx, sqlQuery, moneyU, idU)

	return err
}

func ProductsCostUpdateRow(conn *pgx.Conn, ctx context.Context, costP int, idP int) error { //обновляем строку, id которой совпадает с idP, устанавливая значение cost равным costP

	sqlQuery := `
	UPDATE products
	SET cost = $1
	WHERE productid = $2;
	`

	_, err := conn.Exec(ctx, sqlQuery, costP, idP)

	return err
}

func ProductsAmountUpdateRow(conn *pgx.Conn, ctx context.Context, amountP int, idP int) error { //обновляем строку, id которой совпадает с idP, устанавливая значение amount равным amountP

	sqlQuery := `	
	UPDATE products
	SET amount = $1
	WHERE productid = $2;
	`

	_, err := conn.Exec(ctx, sqlQuery, amountP, idP)

	return err
}
