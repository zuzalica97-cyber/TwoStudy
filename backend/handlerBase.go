package backend

import (
	"fmt"
	"net/http"
	"strconv"
	"study2/market"

	"github.com/gorilla/mux"
)

/*
pattern /base{id}
Mathod GET
Info pattern
*/
func (h *HandlerStruct) HandleGetBase(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["id"]

	titleInt, err := strconv.Atoi(title)

	if err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	fmt.Println(titleInt)

	base, err := h.marketPlase.GetInBase(titleInt)

	if err != nil {
		ErrorDTOmaxiMaker(w, err, market.ErrorBaseNotFound)
		return
	}

	WriteMaker(w, base)
}

/*
pattern /base
Mathod GET
Info ---
*/
func (h *HandlerStruct) HandleGetListBases(w http.ResponseWriter, r *http.Request) {
	list := h.marketPlase.ListBases()

	WriteMaker(w, list)
}
