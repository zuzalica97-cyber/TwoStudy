package market

import "errors"

func (m *Market) NewProdyct(id ProdyctInfo) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if _, ok := m.Prodyct[id.IdP]; ok {
		return ErrorProductAlredyExist
	}

	m.Prodyct[id.IdP] = id
	return nil
}

func (m *Market) GetProdyct(id int) (ProdyctInfo, error) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	prod, ok := m.Prodyct[id]

	if !ok {
		return ProdyctInfo{}, ErrorProductNotFound
	}

	return prod, nil
}

func (m *Market) ListProduct() map[int]ProdyctInfo {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	tmp := make(map[int]ProdyctInfo, len(m.Prodyct))

	for k, v := range m.Prodyct {
		tmp[k] = v
	}

	return tmp
}

func (m *Market) UpCostProduct(money int, id int) (ProdyctInfo, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if money < 0 {
		return ProdyctInfo{}, errors.New("Invalid money")
	}

	prod, ok := m.Prodyct[id]

	if !ok {
		return ProdyctInfo{}, ErrorProductNotFound
	}

	prod.UpCost(money)

	m.Prodyct[id] = prod

	return prod, nil
}

func (m *Market) UpAmountProduct(amount int, id int) (ProdyctInfo, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if amount <= 0 {
		return ProdyctInfo{}, errors.New("Invalid money")
	}

	prod, ok := m.Prodyct[id]

	if !ok {
		return ProdyctInfo{}, ErrorProductNotFound
	}

	prod.UpAmount(amount)

	m.Prodyct[id] = prod

	return prod, nil
}

func (m *Market) DeleteProduct(id int) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	_, ok := m.Prodyct[id]

	if !ok {
		return ErrorProductNotFound
	}

	delete(m.Prodyct, id)
	return nil
}
