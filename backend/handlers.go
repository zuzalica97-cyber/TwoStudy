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

	user, err := h.marketPlase.GetUser(idu)
	if err != nil {
		ErrorDTOmaxiMaker(w, err, market.ErrorUserNotFound)
		return
	}

	prod, err := h.marketPlase.GetProdyct(idp)
	if err != nil {
		ErrorDTOmaxiMaker(w, err, market.ErrorProductNotFound)
		return
	}

	b, err := h.marketPlase.Bay(user, prod, Amount)

	if err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	base, err := json.MarshalIndent(b, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(base); err != nil {
		fmt.Println("fail to write http responce: ", err)
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

	base, err = h.marketPlase.UnBay(base.BaseId)

	if err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	b, err := json.MarshalIndent(base, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("fail to write http responce: ", err)
		return
	}
}
