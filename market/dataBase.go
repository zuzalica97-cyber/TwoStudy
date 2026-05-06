package market

import (
	"fmt"
	"study2/feature_postgres/simple_sql"
)

func (m *Market) AddInBase(base DataBaseInfo) error {

	fmt.Println(base.UserBase.IDU, base.UserBase.ID)

	if err := simple_sql.BaseInsertRow(m.conn, m.ctx, base.UserBase.IDU, base.ProductBase.IDP, base.Ammount, base.Cost, base.Cancelled, base.DataId); err != nil {
		return err
	}
	return nil
}

func (m *Market) GetInBase(id int) (simple_sql.DataBaseModel, error) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	err, base := simple_sql.BaseSelectRow(m.conn, m.ctx, id)

	if err != nil {
		return simple_sql.DataBaseModel{}, fmt.Errorf("Error getting base: %w", err)
	}

	return base, nil
}

func (m *Market) ListBases() map[int]simple_sql.DataBaseModel {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	err, bases := simple_sql.BaseListSelectRow(m.conn, m.ctx)

	if err != nil {
		return nil
	}

	tmp := make(map[int]simple_sql.DataBaseModel, len(bases))

	for _, v := range bases {
		tmp[v.ID] = v
	}

	return tmp
}
