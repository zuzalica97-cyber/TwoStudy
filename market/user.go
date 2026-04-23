package market

import "errors"

func (m *Market) NewUser(id UserInfo) error {

	m.mtx.Lock()
	defer m.mtx.Unlock()

	if _, ok := m.User[id.IdU]; ok {
		return ErrorUserAlredyExist
	}

	m.User[id.IdU] = id
	return nil
}

func (m *Market) GetUser(id int) (UserInfo, error) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	user, ok := m.User[id]

	if !ok {
		return UserInfo{}, ErrorUserNotFound
	}

	return user, nil
}

func (m *Market) ListUser() map[int]UserInfo {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	tmp := make(map[int]UserInfo, len(m.User))

	for k, v := range m.User {
		tmp[k] = v
	}

	return tmp
}

func (m *Market) UpMoneyUser(id int, money int) (UserInfo, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if money <= 0 {
		return UserInfo{}, errors.New("Invalid money valey")
	}

	user, ok := m.User[id]

	if !ok {
		return UserInfo{}, ErrorUserNotFound
	}

	user.UpMoney(money)

	m.User[id] = user

	return user, nil
}

func (m *Market) DeleteUser(id int) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	_, ok := m.User[id]

	if !ok {
		return ErrorUserNotFound
	}

	delete(m.User, id)

	return nil
}
