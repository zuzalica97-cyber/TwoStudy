package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"study2/market"

	"github.com/gorilla/mux"
)

type HandlerStruct struct {
	marketPlase *market.Market
}

func NewHandlerStruct(Market *market.Market) *HandlerStruct {
	return &HandlerStruct{
		marketPlase: Market,
	}
}

/*
pattern /bay
Mathod POST
Info pattern
*/
func (h *HandlerStruct) HandleBay(w http.ResponseWriter, r *http.Request) {
	var bayDTO BayDTO

	if err := json.NewDecoder(r.Body).Decode(&bayDTO); err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	idu := bayDTO.IdUDTO

	idp := bayDTO.IdPDTO

	Amount := bayDTO.AmountDTO

	if _, err := h.marketPlase.GetUser(idu); err != nil {
		ErrorDTOmaxiMaker(w, err, market.ErrorUserNotFound)
		return
	}

	if _, err := h.marketPlase.GetProdyct(idp); err != nil {
		ErrorDTOmaxiMaker(w, err, market.ErrorProductNotFound)
		return
	}

	u, p, err := h.marketPlase.Bay(idu, idp, Amount)

	if err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	b, err := json.MarshalIndent(u, "", "	")
	if err != nil {
		panic(err)
	}
	v, err := json.MarshalIndent(p, "", "	")

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("fail to write http responce: ", err)
		return
	}
	if _, err := w.Write(v); err != nil {
		fmt.Println("fail to write http response2: ", err)
		return
	}
}

/*
pattern /Bay{id}
Mathod DELETE
Info pattern
*/
func (h *HandlerStruct) HandleUnBay(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["id"]

	titleInt, err := strconv.Atoi(title)

	if err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	base, err := h.marketPlase.GetInBase(titleInt)

	if err != nil {
		ErrorDTOmaxiMaker(w, err, market.ErrorBaseNotFound)
		return
	}

	user, prod, err := h.marketPlase.UnBay(base.DataId)

	if err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	b, err := json.MarshalIndent(user, "", "	")
	if err != nil {
		panic(err)
	}
	v, err := json.MarshalIndent(prod, "", "	")

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("fail to write http responce: ", err)
		return
	}
	if _, err := w.Write(v); err != nil {
		fmt.Println("fail to write http response2: ", err)
		return
	}
}
