package main

import (
	"context"
	"fmt"
	"study2/backend"
	"study2/feature_postgres/simple_connection"
	"study2/market"
)

func main() {

	ctx := context.Background()

	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Хованский доволен")

	marketplase := market.NewMarket(ctx, conn)

	httpHandlers := backend.NewHandlerStruct(marketplase)

	httpServer := backend.NewHttpServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server: ", err)
	}
}
