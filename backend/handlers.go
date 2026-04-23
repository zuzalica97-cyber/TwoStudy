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
	var bayDTO ProdyctDTO

	if err := json.NewDecoder(r.Body).Decode(&bayDTO); err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	fmt.Println("fdfdfsddfs")

	idu := bayDTO.Cost

	fmt.Println(idu)

	idp := bayDTO.Amount

	fmt.Println(idp)

	Amount := 1

	fmt.Println(Amount)

	if _, err := h.marketPlase.GetUser(idu); err != nil {
		ErrorDTOmaxiMaker(w, err, market.ErrorUserNotFound)
		return
	}

	fmt.Println("tttt")

	if _, err := h.marketPlase.GetProdyct(idp); err != nil {
		ErrorDTOmaxiMaker(w, err, market.ErrorProductNotFound)
		return
	}

	fmt.Println("pppp")

	u, p, err := h.marketPlase.Bay(idu, idp, Amount)

	fmt.Println("iii")

	if err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	fmt.Println("aaaa")

	WriteMaker(w, u)
	WriteMaker(w, p)
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

	WriteMaker(w, user)
	WriteMaker(w, prod)
}
