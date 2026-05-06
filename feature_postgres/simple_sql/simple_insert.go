package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UsersInsertRow(
	conn *pgx.Conn,
	ctx context.Context,
	Name string,
	Money int,
	UserID int,
) error {

	sqlQuery := `
	INSERT INTO users (username, usermoney, userid) 
	VALUES ($1, $2, $3);
	`

	_, err := conn.Exec(ctx, sqlQuery, Name, Money, UserID)

	return err
}

func ProductsInsertRow(
	conn *pgx.Conn,
	ctx context.Context,
	Name string,
	Description string,
	Cost int,
	Amount int,
	ProductID int,
) error {
	sqlQuery := `
	INSERT INTO products (productname, productdescription, productcost, productamount, productidp) 
	VALUES ($1, $2, $3, $4, $5);
	`
	_, err := conn.Exec(ctx, sqlQuery, Name, Description, Cost, Amount, ProductID)

	return err
}

func BaseInsertRow(
	conn *pgx.Conn,
	ctx context.Context,
	UserID int,
	ProductID int,
	Amount int,
	TotalCost int,
	Canceled bool,
	BaseId int,
) error {

	sqlQuery := `
	INSERT INTO bases (userid, productid, amount, totalcost, canceled, baseid) 
	VALUES ($1, $2, $3, $4, $5, $6);
	`
	_, err := conn.Exec(ctx, sqlQuery, UserID, ProductID, Amount, TotalCost, Canceled, BaseId)

	return err
}
