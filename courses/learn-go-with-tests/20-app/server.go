package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type PlayerServer struct {
	store IPlayerStore
	http.Handler
	mutex sync.Mutex
}

type Player struct {
	Name  string
	Score int
}

func NewServer(store *InMemoryPlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(p.getLeagueTable())
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		p.saveWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) saveWin(w http.ResponseWriter, player string) {
	p.mutex.Lock()
	p.store.SaveWin(player)
	w.WriteHeader(http.StatusAccepted)
	defer p.mutex.Unlock()
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score, ok := p.store.GetPlayerScore(player)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Player not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, score)
}

func (p *PlayerServer) getLeagueTable() []Player {
	return []Player{
		{"Chris", 20},
	}
}
