package market

func (m *Market) AddInBase(base DataBaseInfo) error {

	if _, ok := m.Base[base.DataId]; ok {
		return ErrorBaseAlredyExist
	}

	m.Base[base.DataId] = base
	return nil
}

func (m *Market) GetInBase(id int) (DataBaseInfo, error) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	base, ok := m.Base[id]

	if !ok {
		return DataBaseInfo{}, ErrorBaseNotFound
	}

	return base, nil
}

func (m *Market) ListBases() map[int]DataBaseInfo {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	tmp := make(map[int]DataBaseInfo, len(m.Base))

	for k, v := range m.Base {
		tmp[k] = v
	}

	return tmp
}
