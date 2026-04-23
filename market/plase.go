package market

import (
	"errors"
	"fmt"
	"sync"
)

type Market struct {
	Prodyct map[int]ProdyctInfo
	User    map[int]UserInfo
	Base    map[int]DataBaseInfo
	mtx     sync.RWMutex
}

func NewMarket() *Market {
	return &Market{
		Prodyct: make(map[int]ProdyctInfo),
		User:    make(map[int]UserInfo),
		Base:    make(map[int]DataBaseInfo),
	}
}

func (m *Market) Bay(idU int, idP int, amount int) (UserInfo, ProdyctInfo, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	fmt.Println(1)

	prod, ok := m.Prodyct[idP]

	if !ok {
		return UserInfo{}, ProdyctInfo{}, ErrorProductNotFound
	}

	user, ok := m.User[idU]

	if !ok {
		return UserInfo{}, ProdyctInfo{}, ErrorUserNotFound
	}

	fmt.Println(2)

	if user.Money-prod.Cost*amount < 0 {
		return UserInfo{}, ProdyctInfo{}, errors.New("Not enough money")
	}

	fmt.Println(3)

	if prod.Amount-amount < 0 {
		return UserInfo{}, ProdyctInfo{}, errors.New("Not enough product")
	}

	fmt.Println(3)

	cost := prod.Cost

	user.BayUser(cost, amount)

	fmt.Println(4)

	m.User[idU] = user

	prod.BayProduct(amount)

	fmt.Println(5)

	m.Prodyct[idP] = prod

	MarketBase := MakeDataBase(user.IdU, prod.IdP, prod.Cost, amount)

	fmt.Println(6)

	if err := m.AddInBase(MarketBase); err != nil {
		return UserInfo{}, ProdyctInfo{}, err
	}

	fmt.Println(7)

	return user, prod, nil
}

func (m *Market) UnBay(idB int) (UserInfo, ProdyctInfo, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	base, ok := m.Base[idB]

	if !ok {
		return UserInfo{}, ProdyctInfo{}, ErrorBaseNotFound
	}

	user, ok := m.User[base.UserId]

	if !ok {
		return UserInfo{}, ProdyctInfo{}, ErrorUserNotFound
	}
	prod, ok := m.Prodyct[base.ProductId]

	if !ok {
		return UserInfo{}, ProdyctInfo{}, ErrorProductNotFound
	}

	user.UpMoney(base.BayCost)
	prod.UpAmount(base.BayAmount)

	m.User[base.UserId] = user
	m.Prodyct[base.ProductId] = prod

	base.Cancel()

	return user, prod, nil
}
