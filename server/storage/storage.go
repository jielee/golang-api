package storage

type StoreEr interface {
	GetOne(key interface{}) interface{}
	Create(key interface{}, res interface{}) *interface{}
	Delete(key interface{}) bool
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

func (s *MemoryStorage) Delete(key interface{}) bool{
	if(s.GetOne(key) != nil){
		delete(s.Storage, key)
		return true
	}
	return false
}
