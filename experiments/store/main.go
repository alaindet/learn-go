package main

import (
	"fmt"
)

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

func main() {
	s := NewStore(map[string]interface{}{
		"a": 111,
		"b": 222,
		"c": 333,
	})

	val, err := s.MustGet("zzz")

	if err != nil {
		fmt.Printf("No value for key %q in store\n", "zzz")
	} else {
		fmt.Println("%s: %+v\n", "zzz", val)
	}

	fmt.Println(
		s.Get("ciao"),
		s.Get("a"),
	)
}
