package market

import (
	"context"
	"errors"
	"fmt"
	"study2/feature_postgres/simple_sql"
	"sync"

	"github.com/jackc/pgx/v5"
)

type Market struct {
	ctx  context.Context
	conn *pgx.Conn
	mtx  sync.RWMutex
}

func NewMarket(Ctx context.Context, Conn *pgx.Conn) *Market {
	return &Market{
		ctx:  Ctx,
		conn: Conn,
	}
}

func (m *Market) Bay(user simple_sql.UserModel, prod simple_sql.ProductModel, amount int) (DataBaseInfo, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if user.Money-prod.Cost*amount < 0 {
		return DataBaseInfo{}, errors.New("Not enough money")
	}

	if prod.Amount-amount < 0 {
		return DataBaseInfo{}, errors.New("Not enough product")
	}

	user.Money -= prod.Cost * amount
	prod.Amount -= amount

	fmt.Println(user.Money, prod.Amount)

	err := simple_sql.UserMoneyUpdateRow(m.conn, m.ctx, user.Money, user.IDU)
	if err != nil {
		return DataBaseInfo{}, err
	}

	err = simple_sql.ProductsAmountUpdateRow(m.conn, m.ctx, prod.Amount, prod.IDP)
	if err != nil {
		return DataBaseInfo{}, err
	}

	MarketBase := MakeDataBase(&user, &prod, amount)

	if err := m.AddInBase(MarketBase); err != nil {
		return DataBaseInfo{}, err
	}

	return MarketBase, nil
}

func (m *Market) UnBay(idB int) (simple_sql.DataBaseModel, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	err, base := simple_sql.BaseSelectRow(m.conn, m.ctx, idB)

	if err != nil {
		return simple_sql.DataBaseModel{}, err
	}

	if base.Canceled {
		return simple_sql.DataBaseModel{}, errors.New("Base already canceled")
	}

	err, user := simple_sql.UserSelectRow(m.conn, m.ctx, base.UserID)

	if err != nil {
		return simple_sql.DataBaseModel{}, err
	}

	err, prod := simple_sql.ProdSelectRow(m.conn, m.ctx, base.ProductID)

	if err != nil {
		return simple_sql.DataBaseModel{}, err
	}

	user.Money += prod.Cost * base.Amount

	prod.Amount += base.Amount

	if err := simple_sql.UserMoneyUpdateRow(m.conn, m.ctx, user.Money, user.IDU); err != nil {
		return simple_sql.DataBaseModel{}, err
	}

	if err := simple_sql.ProductsAmountUpdateRow(m.conn, m.ctx, prod.Amount, prod.IDP); err != nil {
		return simple_sql.DataBaseModel{}, err
	}

	if err := simple_sql.BaseCancelUpdateRow(m.conn, m.ctx, true, idB); err != nil {
		return simple_sql.DataBaseModel{}, err
	}

	return base, err
}
