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
	SET productcost = $1
	WHERE productidp = $2;
	`

	_, err := conn.Exec(ctx, sqlQuery, costP, idP)

	return err
}

func ProductsAmountUpdateRow(conn *pgx.Conn, ctx context.Context, amountP int, idP int) error { //обновляем строку, id которой совпадает с idP, устанавливая значение amount равным amountP

	sqlQuery := `	
	UPDATE products
	SET productamount = $1
	WHERE productidp = $2;
	`

	_, err := conn.Exec(ctx, sqlQuery, amountP, idP)

	return err
}

func BaseCancelUpdateRow(conn *pgx.Conn, ctx context.Context, canceled bool, idB int) error { //обновляем строку, id которой совпадает с idB, устанавливая значение canceled равным canceled

	sqlQuery := `
	UPDATE bases
	SET canceled = $1
	WHERE baseid = $2;
	`

	_, err := conn.Exec(ctx, sqlQuery, canceled, idB)

	return err
}
