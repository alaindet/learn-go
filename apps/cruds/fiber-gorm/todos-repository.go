package main

import (
	"database/sql"
)

func createTodo(db *sql.DB, dto *CreateTodoDto) (*Todo, error) {

	stmt, err := db.Prepare("INSERT INTO todos (name) VALUES (?)")

	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(dto.Name)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	todo := &Todo{
		Id:     id,
		Name:   dto.Name,
		IsDone: false,
	}

	return todo, nil
}

// TODO: Use prepared statements?
func getTodoById(db *sql.DB, id int64) (*Todo, error) {

	todo := &Todo{}

	row := db.QueryRow("SELECT * FROM todos WHERE id = ?", id)

	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(
		&todo.Id,
		&todo.Name,
		&todo.IsDone,
	)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

// TODO: Use prepared statements?
func getTodos(db *sql.DB) ([]Todo, error) {

	rows, err := db.Query("SELECT * FROM todos")

	if err != nil {
		return nil, err
	}

	todos := []Todo{} // TODO: Set size accordingly?

	for rows.Next() {
		t := Todo{}
		err := rows.Scan(&t.Id, &t.Name, &t.IsDone)

		if err != nil {
			return todos, err
		}

		todos = append(todos, t)
	}

	return todos, nil
}

func updateTodo(db *sql.DB, dto *UpdateTodoDto) (*Todo, error) {

	todo, err := getTodoById(db, dto.Id)

	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare("UPDATE todos SET name = ?, is_done = ? WHERE id = ?")

	if err != nil {
		return todo, err
	}

	result, err := stmt.Exec(dto.Name, dto.IsDone, dto.Id)
	_ = result

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func deleteTodo(db *sql.DB, dto *DeleteTodoDto) (*Todo, error) {

	todo, err := getTodoById(db, dto.Id)

	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare("DELETE FROM todos WHERE id = ?")

	if err != nil {
		return todo, err
	}

	result, err := stmt.Exec(dto.Id)
	_ = result

	if err != nil {
		return todo, err
	}

	// TODO: Check count?

	return todo, nil
}
