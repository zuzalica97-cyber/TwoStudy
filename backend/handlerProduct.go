package backend

import (
	"encoding/json"
	"net/http"
	"strconv"
	"study2/market"

	"github.com/gorilla/mux"
)

/*
pattern /prod
Mathod POST
Info JSON in reqwest body
*/
func (h *HandlerStruct) HandleNewProduct(w http.ResponseWriter, r *http.Request) {
	var prodDTO ProdyctDTO

	if err := json.NewDecoder(r.Body).Decode(&prodDTO); err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	if err := prodDTO.ValidateForCreateProduct(); err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	prod := market.MakeProduct(prodDTO.Name, prodDTO.Description, prodDTO.Cost, prodDTO.Amount)

	if err := h.marketPlase.NewProdyct(prod); err != nil {
		ErrorDTOmaxiMaker(w, err, market.ErrorProductAlredyExist)
		return
	}

	WriteMaker(w, prod)
}

/*
pattern /prod{id}
Mathod GET
Info pattern
*/
func (h *HandlerStruct) HandleGetProduct(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["id"]

	titleInt, err := strconv.Atoi(title)

	if err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	prod, err := h.marketPlase.GetProdyct(titleInt)

	if err != nil {
		ErrorDTOmaxiMaker(w, err, market.ErrorProductNotFound)
		return
	}

	WriteMaker(w, prod)
}

/*
pattern /prod
Mathod GET
Info ---
*/
func (h *HandlerStruct) HandleGetListProducts(w http.ResponseWriter, r *http.Request) {
	list := h.marketPlase.ListProduct()

	WriteMaker(w, list)
}

/*
pattern /prod{id}?cost=true
Mathod PATCH
Info pattern + JSON
*/
func (h *HandlerStruct) HandleUpCostProduct(w http.ResponseWriter, r *http.Request) {
	var costDTO CostDTO

	title := mux.Vars(r)["id"]

	titleInt, err := strconv.Atoi(title)

	if err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&costDTO); err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	if err := costDTO.ValidateForCreateCost(); err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	prod, err := h.marketPlase.UpCostProduct(titleInt, costDTO.NewCost)

	if err != nil {
		ErrorDTOmaxiMaker(w, err, market.ErrorProductNotFound)
		return
	}

	WriteMaker(w, prod)
}

/*
pattern /prod{id}?cost=false
Mathod PATCH
Info pattern + JSON
*/
func (h *HandlerStruct) HandleUpAmountProduct(w http.ResponseWriter, r *http.Request) {
	var amountDTO AmountDTO

	title := mux.Vars(r)["id"]

	titleInt, err := strconv.Atoi(title)

	if err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&amountDTO); err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	if err := amountDTO.ValidateForCreateAmount(); err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	prod, err := h.marketPlase.UpAmountProduct(titleInt, amountDTO.NewAmount)

	if err != nil {
		ErrorDTOmaxiMaker(w, err, market.ErrorProductNotFound)
		return
	}

	WriteMaker(w, prod)
}

/*
pattern /prod{id}
Mathod DELETE
Info pattern
*/
func (h *HandlerStruct) HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["id"]

	titleInt, err := strconv.Atoi(title)

	if err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	if err := h.marketPlase.DeleteProduct(titleInt); err != nil {
		ErrorDTOmaxiMaker(w, err, market.ErrorProductNotFound)
	}
	w.WriteHeader(http.StatusNoContent)
}
