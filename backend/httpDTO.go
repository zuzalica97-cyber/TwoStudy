package backend

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type ProdyctDTO struct {
	Name        string
	Description string
	Cost        int
	Amount      int
	aaaa        int
}

type UserDTO struct {
	Name  string
	Money int
}

type ErrorDTO struct { //Чтобы можно было возращять информацию об ошибке
	Message string
	Time    time.Time
}

type BayDTO struct {
	IdUDTO    int
	IdPDTO    int
	AmountDTO int
}

func (p ProdyctDTO) ValidateForCreateProduct() error { //проверяем если одно из знужных нам значений отсутсвует то возвращяем значения
	if p.Name == "" {
		return errors.New("title is empty")
	}

	if p.Description == "" {
		return errors.New("description is empty")
	}

	if p.Cost == 0 {
		return errors.New("cost is empty")
	}

	if p.Amount == 0 {
		return errors.New("amount is empty")
	}
	return nil
}

func (u UserDTO) ValidateForCreateUser() error { //проверяем если одно из знужных нам значений отсутсвует то возвращяем значения
	if u.Name == "" {
		return errors.New("title is empty")
	}

	if u.Money == 0 {
		return errors.New("money is empty")
	}
	return nil
}

func (b BayDTO) ValidateForBay() error {
	if b.IdUDTO == 0 {
		return errors.New("user is empty")
	}
	if b.IdPDTO == 0 {
		return errors.New("prod is empty")
	}
	if b.AmountDTO == 0 {
		return errors.New("amount is empty")
	}
	return nil
}

func (e *ErrorDTO) ToString() string { //приобразуем данные об ошибки для красивого вывода
	b, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		panic(err)
	}

	return string(b) //возвращяем нашу ошибку в виде строки
}

type CostDTO struct {
	NewCost int
}

type AmountDTO struct {
	NewAmount int
}

type MoneyDTO struct {
	NewMoney int
}

func (c CostDTO) ValidateForCreateCost() error {
	if c.NewCost == 0 {
		return errors.New("New cost is empty")
	}
	return nil
}

func (c AmountDTO) ValidateForCreateAmount() error {
	if c.NewAmount == 0 {
		return errors.New("New amount is empty")
	}
	return nil
}

func (c MoneyDTO) ValidateForCreateMoney() error {
	if c.NewMoney == 0 {
		return errors.New("New money is empty")
	}
	return nil
}

func ErrorDTOmaker(err error, w http.ResponseWriter) {

	errDTO := ErrorDTO{
		Message: err.Error(),
		Time:    time.Now(),
	}

	http.Error(w, errDTO.ToString(), http.StatusBadRequest)
}

func ErrorDTOmaxiMaker(w http.ResponseWriter, err error, marketEror error) {
	errDTO := ErrorDTO{
		Message: err.Error(),
		Time:    time.Now(),
	}

	if errors.Is(err, marketEror) { //проверяем кокого типа эта ошибка тут праверка на 409
		http.Error(w, errDTO.ToString(), http.StatusConflict)

	} else {
		http.Error(w, errDTO.ToString(), http.StatusConflict) // если нет то 500
	}
}

func WriteMaker(w http.ResponseWriter, name any) {
	b, err := json.MarshalIndent(name, "", "	")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("fail to write http responce: ", err)
		return
	}
}
