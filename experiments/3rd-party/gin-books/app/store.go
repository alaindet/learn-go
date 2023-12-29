package app

import "fmt"

type Store struct {
	data map[string]interface{}
}

func NewStore(initialData ...map[string]interface{}) *Store {

	var data map[string]interface{}

	if len(initialData) > 0 {
		data = initialData[0]
	} else {
		data = make(map[string]interface{})
	}

	return &Store{
		data: data,
	}
}

func (s *Store) Get(name string) interface{} {
	item, ok := s.data[name]

	if !ok {
		return nil
	}

	return item
}

func (s *Store) MustGet(name string) (interface{}, error) {
	item, ok := s.data[name]

	if !ok {
		return nil, fmt.Errorf("No data with name %q in store", name)
	}

	return item, nil
}

func (s *Store) Set(name string, item interface{}) {
	s.data[name] = item
}
