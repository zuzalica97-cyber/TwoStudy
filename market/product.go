package market

import (
	"errors"
	"fmt"
	"study2/feature_postgres/simple_sql"
)

func (m *Market) NewProdyct(p ProdyctInfo) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if err := simple_sql.ProductsInsertRow(m.conn, m.ctx, p.Name, p.Description, p.Cost, p.Amount, p.IdP); err != nil {
		return err
	}

	return nil
}

func (m *Market) GetProdyct(id int) (simple_sql.ProductModel, error) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	err, prod := simple_sql.ProdSelectRow(m.conn, m.ctx, id)

	if err != nil {
		return simple_sql.ProductModel{}, ErrorProductNotFound
	}

	return prod, nil
}

func (m *Market) ListProduct() map[int]simple_sql.ProductModel {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	err, prods := simple_sql.ProdListSelectRow(m.conn, m.ctx)

	if err != nil {
		return nil
	}

	tmp := make(map[int]simple_sql.ProductModel, len(prods))

	for _, v := range prods {
		tmp[v.ID] = v
	}

	return tmp
}

func (m *Market) UpCostProduct(money int, id int) (simple_sql.ProductModel, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if money <= 0 {
		return simple_sql.ProductModel{}, errors.New("Invalid money")
	}

	err, prod := simple_sql.ProdSelectRow(m.conn, m.ctx, id)

	if err != nil {
		return simple_sql.ProductModel{}, ErrorProductNotFound
	}

	prod.Cost += money

	if err := simple_sql.ProductsCostUpdateRow(m.conn, m.ctx, prod.Cost, id); err != nil {
		return simple_sql.ProductModel{}, err
	}

	return prod, nil
}

func (m *Market) UpAmountProduct(amount int, id int) (simple_sql.ProductModel, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if amount <= 0 {
		return simple_sql.ProductModel{}, errors.New("Invalid amount")
	}

	err, prod := simple_sql.ProdSelectRow(m.conn, m.ctx, id)

	if err != nil {
		return simple_sql.ProductModel{}, ErrorProductNotFound
	}

	prod.Amount += amount

	if err := simple_sql.ProductsAmountUpdateRow(m.conn, m.ctx, prod.Amount, id); err != nil {
		return simple_sql.ProductModel{}, err
	}

	return prod, nil
}

func (m *Market) DeleteProduct(id []int) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	fmt.Println(id)

	if err := simple_sql.ProductDeleteRow(m.conn, m.ctx, id); err != nil {
		return ErrorProductNotFound
	}

	return nil
}
