package main

import (
	"fmt"
	"study2/backend"
	"study2/feature_postgres/simple_connection"
	"study2/market"
)

func main() {

	simple_connection.CheckConnection()

	marketplase := market.NewMarket()

	httpHandlers := backend.NewHandlerStruct(marketplase)

	httpServer := backend.NewHttpServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server: ", err)
	}
}
