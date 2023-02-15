package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	Name        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []Todo

func NewTodos() *Todos {
	return &Todos{}
}

func (t *Todos) Add(name string) error {
	*t = append(*t, Todo{
		Name:        name,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	})

	return nil
}

func (t *Todos) Get(i int) (*Todo, error) {

	if err := t.checkIndex(i); err != nil {
		return nil, err
	}

	return &(*t)[i], nil
}

func (t *Todos) Complete(i int) error {

	if err := t.checkIndex(i); err != nil {
		return err
	}

	todos := *t
	todos[i].Done = true
	todos[i].CompletedAt = time.Now()
	return nil
}

func (t *Todos) Delete(i int) error {

	if err := t.checkIndex(i); err != nil {
		return err
	}

	todos := *t
	*t = append(todos[:i-1], todos[i:]...)
	return nil
}

func (t *Todos) SaveToStorage(filename string) error {
	js, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, js, 0644)
}

func (t *Todos) FetchFromStorage(filename string) error {
	file, err := os.ReadFile(filename)

	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	if err != nil {
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, t)
}

func (t *Todos) checkIndex(i int) error {

	if i < 0 || i > (len(*t)-1) {
		return fmt.Errorf("todo #%d does not exist", i)
	}

	return nil
}
