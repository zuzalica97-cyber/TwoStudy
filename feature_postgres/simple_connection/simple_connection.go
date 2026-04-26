package simple_connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:Ww2425wW4245G@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Подключеник к БАЗЕ прошло успешно. Хованский доволен")
}
