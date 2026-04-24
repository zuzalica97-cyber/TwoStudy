package market

import (
	"errors"
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

	prod, ok := m.Prodyct[idP]

	if !ok {
		return UserInfo{}, ProdyctInfo{}, ErrorProductNotFound
	}

	user, ok := m.User[idU]

	if !ok {
		return UserInfo{}, ProdyctInfo{}, ErrorUserNotFound
	}

	if user.Money-prod.Cost*amount < 0 {
		return UserInfo{}, ProdyctInfo{}, errors.New("Not enough money")
	}

	if prod.Amount-amount < 0 {
		return UserInfo{}, ProdyctInfo{}, errors.New("Not enough product")
	}

	cost := prod.Cost

	user.BayUser(cost, amount)

	m.User[idU] = user

	prod.BayProduct(amount)

	m.Prodyct[idP] = prod

	MarketBase := MakeDataBase(user.IdU, prod.IdP, prod.Cost, amount)

	if err := m.AddInBase(MarketBase); err != nil {
		return UserInfo{}, ProdyctInfo{}, err
	}

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
	m.Base[base.DataId] = base

	return user, prod, nil
}
