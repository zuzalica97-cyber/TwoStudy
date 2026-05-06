package market

import (
	"errors"
	"math/rand"
	"study2/feature_postgres/simple_sql"
)

type ProdyctInfo struct {
	IdP         int
	Name        string
	Description string
	Cost        int
	Amount      int
}

type UserInfo struct {
	IdU   int
	Name  string
	Money int
}

type DataBaseInfo struct {
	ProductBase simple_sql.ProductModel
	UserBase    simple_sql.UserModel
	Ammount     int
	Cost        int
	DataId      int
	Cancelled   bool
}

func MakeProduct(name string, description string, cost int, amount int) ProdyctInfo {

	if cost < 0 {
		return ProdyctInfo{}
	}

	if amount < 0 {
		return ProdyctInfo{}
	}

	return ProdyctInfo{
		IdP:         rand.Intn(10000),
		Name:        name,
		Description: description,
		Cost:        cost,
		Amount:      amount,
	}
}

func MakeUser(name string, money int) (error, UserInfo) {

	if money < 0 {
		return errors.New("Invalod money valeuy"), UserInfo{}
	}

	return nil, UserInfo{
		IdU:   rand.Intn(10000),
		Name:  name,
		Money: money,
	}
}

func MakeDataBase(user *simple_sql.UserModel, prod *simple_sql.ProductModel, amount int) DataBaseInfo {

	return DataBaseInfo{
		ProductBase: *prod,
		UserBase:    *user,
		Ammount:     amount,
		DataId:      rand.Intn(10000),
		Cost:        prod.Cost * amount,
		Cancelled:   false,
	}
}
