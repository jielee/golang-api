package storage

type StoreEr interface {
	GetOne(key interface{}) interface{}
	Create(key interface{}, res interface{}) *interface{}
}

type MemoryStorage struct {
	Storage map[interface{}]interface{}
}

func (s *MemoryStorage) GetOne(key interface{}) interface{}{
	return s.Storage[key]
}

func (s *MemoryStorage) Create(key interface{}, res interface{}) *interface{}{
	s.Storage[key] = res
	return &res
}
