package backend

import (
	"encoding/json"
	"net/http"
	"strconv"
	"study2/market"

	"github.com/gorilla/mux"
)

/*
pattern /user
Mathod POST
Info JSON in reqwest body
*/
func (h *HandlerStruct) HandleNewUser(w http.ResponseWriter, r *http.Request) {
	var userDTO UserDTO

	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	if err := userDTO.ValidateForCreateUser(); err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	user := market.MakeUser(userDTO.Name, userDTO.Money)

	if err := h.marketPlase.NewUser(user); err != nil {
		ErrorDTOmaxiMaker(w, err, market.ErrorUserAlredyExist)
		return
	}

	WriteMaker(w, user)
}

/*
pattern /user{id}
Mathod GET
Info patteern
*/
func (h *HandlerStruct) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["id"]

	titleInt, err := strconv.Atoi(title)

	if err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	user, err := h.marketPlase.GetUser(titleInt)

	if err != nil {
		ErrorDTOmaxiMaker(w, err, market.ErrorUserNotFound)
		return
	}

	WriteMaker(w, user)
}

/*
pattern /user
Mathod GET
Info ---
*/
func (h *HandlerStruct) HandleGetListUsers(w http.ResponseWriter, r *http.Request) {
	list := h.marketPlase.ListUser()

	WriteMaker(w, list)
}

/*
pattern /user/{id}
Mathod PATCH
Info pattern + JSON
*/
func (h *HandlerStruct) HandleUpMoneyUser(w http.ResponseWriter, r *http.Request) {
	var moneyDTO MoneyDTO

	title := mux.Vars(r)["id"]

	titleInt, err := strconv.Atoi(title)

	if err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&moneyDTO); err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	if err := moneyDTO.ValidateForCreateMoney(); err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	user, err := h.marketPlase.UpMoneyUser(titleInt, moneyDTO.NewMoney)

	if err != nil {
		ErrorDTOmaxiMaker(w, err, market.ErrorUserNotFound)
		return
	}

	WriteMaker(w, user)
}

/*
pattern /user{id}
Mathod DELETE
Info pattern
*/
func (h *HandlerStruct) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["id"]

	titleInt, err := strconv.Atoi(title)

	if err != nil {
		ErrorDTOmaker(err, w)
		return
	}

	if err := h.marketPlase.DeleteUser(titleInt); err != nil {
		ErrorDTOmaxiMaker(w, err, market.ErrorUserNotFound)
	}
	w.WriteHeader(http.StatusNoContent)
}
