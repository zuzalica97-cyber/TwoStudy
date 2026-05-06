package market

import (
	"errors"
	"study2/feature_postgres/simple_sql"
)

func (m *Market) NewUser(user UserInfo) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if err := simple_sql.UsersInsertRow(m.conn, m.ctx, user.Name, user.Money, user.IdU); err != nil {
		return err
	}

	return nil
}

func (m *Market) GetUser(id int) (simple_sql.UserModel, error) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	err, user := simple_sql.UserSelectRow(m.conn, m.ctx, id)

	if err != nil {
		return simple_sql.UserModel{}, err
	}

	return user, nil
}

func (m *Market) ListUser() map[int]simple_sql.UserModel {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	err, users := simple_sql.UserListSelectRow(m.conn, m.ctx)

	if err != nil {
		return nil
	}

	tmp := make(map[int]simple_sql.UserModel, len(users))

	for _, v := range users {
		tmp[v.ID] = v
	}

	return tmp
}

func (m *Market) UpMoneyUser(id int, money int) (simple_sql.UserModel, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if money <= 0 {
		return simple_sql.UserModel{}, errors.New("Invalid money value")
	}

	err, user := simple_sql.UserSelectRow(m.conn, m.ctx, id)

	if err != nil {
		return simple_sql.UserModel{}, ErrorUserNotFound
	}

	user.Money += money

	if err := simple_sql.UserMoneyUpdateRow(m.conn, m.ctx, user.Money, id); err != nil {
		return simple_sql.UserModel{}, err
	}

	return user, nil
}

func (m *Market) DeleteUser(id []int) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if err := simple_sql.UserDeleteRow(m.conn, m.ctx, id); err != nil {
		return ErrorUserNotFound
	}

	return nil
}
