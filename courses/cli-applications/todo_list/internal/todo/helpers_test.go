package todo_test

import (
	"testing"

	"todo_list/internal/todo"
)

func assertTodoNameAtIndex(t testing.TB, todos *todo.Todos, index int, expectedName string) {
	t.Helper()
	todo, err := todos.Get(index)

	if err != nil {
		t.Error(err)
	}

	if todo.Name != expectedName {
		t.Errorf("Expected %q, got %q instead.", expectedName, todo.Name)
	}
}

func assertTaskIsDoneAtIndex(t testing.TB, todos *todo.Todos, index int) {
	t.Helper()
	todo, err := todos.Get(index)

	if err != nil {
		t.Error(err)
	}

	if !todo.Done {
		t.Errorf("Task should be completed")
	}
}

func assertTaskIsNotDoneAtIndex(t testing.TB, todos *todo.Todos, index int) {
	t.Helper()
	todo, err := todos.Get(index)

	if err != nil {
		t.Error(err)
	}

	if todo.Done {
		t.Errorf("Task should not be completed")
	}
}

func assertTasksCount(t testing.TB, todos *todo.Todos, expectedCount int) {
	t.Helper()
	todosCount := len(*todos)
	if todosCount != expectedCount {
		t.Errorf("Expected list length %d, got %d instead.", todosCount, expectedCount)
	}
}
