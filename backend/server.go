package backend

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	httpHandlers *HandlerStruct
}

func NewHttpServer(httpHandler *HandlerStruct) *HttpServer {
	return &HttpServer{
		httpHandlers: httpHandler,
	}
}

func (s *HttpServer) StartServer() error {
	router := mux.NewRouter()

	//USERS

	router.Path("/user").Methods("POST").HandlerFunc(s.httpHandlers.HandleNewUser)

	router.Path("/user/{id}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetUser)

	router.Path("/user").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetListUsers)

	router.Path("/user/{id}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandleUpMoneyUser)

	router.Path("/user/{id}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandleDeleteUser)

	//PRODUCT

	router.Path("/prod").Methods("POST").HandlerFunc(s.httpHandlers.HandleNewProduct)

	router.Path("/prod/{id}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetProduct)

	router.Path("/prod").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetListProducts)

	router.Path("/prod/{id}").Methods("PATCH").Queries("cost", "true").HandlerFunc(s.httpHandlers.HandleUpCostProduct)

	router.Path("/prod/{id}").Methods("PATCH").Queries("cost", "false").HandlerFunc(s.httpHandlers.HandleUpAmountProduct)

	router.Path("/prod/{id}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandleDeleteProduct)

	//BASE

	router.Path("/base/{id}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetBase)

	router.Path("/base").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetListBases)

	//BAY

	router.Path("/bay").Methods("POST").HandlerFunc(s.httpHandlers.HandleBay)

	router.Path("/bay/{id}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandleUnBay)

	if err := http.ListenAndServe(":9021", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}

	return nil
}
