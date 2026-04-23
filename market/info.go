package market

import "math/rand"

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
	UserId    int
	ProductId int
	BayCost   int
	BayAmount int
	DataId    int
	Cancelled bool
}

func MakeProduct(name string, description string, cost int, amount int) ProdyctInfo {

	if cost < 0 {
		return ProdyctInfo{}
	}

	if amount < 0 {
		return ProdyctInfo{}
	}

	return ProdyctInfo{
		IdP:         rand.Intn(1000),
		Name:        name,
		Description: description,
		Cost:        cost,
		Amount:      amount,
	}
}

func MakeUser(name string, money int) UserInfo {

	if money < 0 {
		return UserInfo{}
	}

	return UserInfo{
		IdU:   rand.Intn(1000),
		Name:  name,
		Money: money,
	}
}

func MakeDataBase(userId int, productId int, Cost int, bayAmount int) DataBaseInfo {

	bayCost := Cost * bayAmount

	return DataBaseInfo{
		UserId:    userId,
		ProductId: productId,
		BayCost:   bayCost,
		BayAmount: bayAmount,
		DataId:    rand.Intn(1000),
		Cancelled: false,
	}
}

func (p *ProdyctInfo) UpCost(money int) {
	p.Cost += money
}

func (p *ProdyctInfo) UpAmount(amount int) {
	p.Amount += amount
}

func (u *UserInfo) UpMoney(money int) {
	u.Money += money
}

func (u *UserInfo) BayUser(cost int, amount int) {
	u.Money -= cost * amount
}

func (p *ProdyctInfo) BayProduct(amount int) {
	p.Amount -= amount
}

func (b *DataBaseInfo) Cancel() {
	b.Cancelled = true
}
