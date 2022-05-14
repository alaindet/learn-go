package main

type IPlayerStore interface {
	SaveWin(name string)
	GetPlayerScore(name string) (int, bool)
}

type InMemoryPlayerStore struct {
	scores map[string]int
}

func NewInMemoryStore(args ...map[string]int) *InMemoryPlayerStore {

	var scores map[string]int

	if len(args) == 1 && args[0] != nil {
		scores = args[0]
	} else {
		scores = make(map[string]int, 10)
	}

	return &InMemoryPlayerStore{scores}
}

func (s *InMemoryPlayerStore) SaveWin(name string) {
	s.scores[name] += 1
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) (int, bool) {
	val, ok := s.scores[name]
	return val, ok
}
