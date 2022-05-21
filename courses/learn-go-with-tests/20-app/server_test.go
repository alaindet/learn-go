package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, bool) {
	score, ok := s.scores[name]
	return score, ok
}

func (s *StubPlayerStore) SaveWin(name string) {
	s.scores[name] += 1
}

func TestGETPlayers(t *testing.T) {

	scores := map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	}

	store := NewInMemoryStore(scores)
	server := NewServer(store)

	for name, score := range scores {
		testName := fmt.Sprintf("returns %s's score", name)
		t.Run(testName, func(t *testing.T) {
			request := newGetScoreRequest(name)
			response := httptest.NewRecorder()
			server.ServeHTTP(response, request)
			result := response.Body.String()
			expected := strconv.Itoa(score)
			assertResponseBody(t, result, expected)
			assertStatus(t, response.Code, 200)
		})
	}

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, 404)
	})
}

func TestStoreWins(t *testing.T) {

	store := NewInMemoryStore()
	server := NewServer(store)

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)
		score, _ := store.GetPlayerScore(player)
		assertScore(t, score, 1)
	})
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryStore()
	server := NewServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}

func TestLeague(t *testing.T) {
	store := NewInMemoryStore()
	server := NewServer(store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		var result []Player
		err := json.NewDecoder(response.Body).Decode(&result)
		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
		}
		assertStatus(t, response.Code, http.StatusOK)
	})
}

func newGetScoreRequest(name string) *http.Request {
	url := fmt.Sprintf("/players/%s", name)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	url := fmt.Sprintf("/players/%s", name)
	req, _ := http.NewRequest(http.MethodPost, url, nil)
	return req
}

func assertResponseBody(t testing.TB, result, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("response body is wrong, result %q expected %q", result, expected)
	}
}

func assertStatus(t testing.TB, result, expected int) {
	t.Helper()
	if result != expected {
		t.Errorf("got HTTP status %d expected %d", result, expected)
	}
}

func assertScore(t testing.TB, result, expected int) {
	t.Helper()
	if result != expected {
		t.Errorf("score is %d expected %d", result, expected)
	}
}
