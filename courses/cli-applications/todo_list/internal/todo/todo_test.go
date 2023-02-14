package todo_test

import (
	"os"
	"testing"

	"todo_list/internal/todo"
)

func TestAdd(t *testing.T) {
	todos := todo.NewTodos()
	todoName := "New Todo"
	todos.Add(todoName)
	assertTodoNameAtIndex(t, todos, 0, todoName)
}

func TestComplete(t *testing.T) {
	todos := todo.NewTodos()
	todoName := "New Todo"
	todos.Add(todoName)
	assertTodoNameAtIndex(t, todos, 0, todoName)
	assertTaskIsNotDoneAtIndex(t, todos, 0)
	todos.Complete(0)
	assertTaskIsDoneAtIndex(t, todos, 0)
}

func TestDelete(t *testing.T) {
	todos := todo.NewTodos()
	newTodos := []string{"New Todo 1", "New Todo 2", "New Todo 3"}
	for _, v := range newTodos {
		todos.Add(v)
	}
	assertTodoNameAtIndex(t, todos, 0, newTodos[0])
	assertTodoNameAtIndex(t, todos, 1, newTodos[1])
	assertTodoNameAtIndex(t, todos, 2, newTodos[2])
	todos.Delete(1)
	assertTasksCount(t, todos, 2)
	assertTodoNameAtIndex(t, todos, 1, newTodos[2])
}

func TestInteractWithFileSystem(t *testing.T) {
	todos1 := todo.NewTodos()
	todos2 := todo.NewTodos()
	todoName := "New Todo"
	todos1.Add(todoName)
	assertTodoNameAtIndex(t, todos1, 0, todoName)
	tempFile, err := os.CreateTemp("", "")

	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}

	tempFileName := tempFile.Name()
	defer os.Remove(tempFileName)

	err = todos1.SaveToStorage(tempFileName)

	if err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	err = todos2.FetchFromStorage(tempFileName)

	if err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	todo1, _ := todos1.Get(0)
	todo2, _ := todos2.Get(0)

	if todo1.Name != todo2.Name {
		t.Errorf("Task %q should match %q task.", todo1.Name, todo2.Name)
	}
}
