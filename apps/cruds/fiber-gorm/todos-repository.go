package main

import (
	"database/sql"
)

type todosRepository struct {
	db *sql.DB
}

func (r *todosRepository) create(dto *CreateTodoDto) (*Todo, error) {
	stmt, err := r.db.Prepare("INSERT INTO todos (name) VALUES (?)")

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
func (r *todosRepository) getById(id int64) (*Todo, error) {

	todo := &Todo{}

	row := r.db.QueryRow("SELECT * FROM todos WHERE id = ?", id)

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
func (r *todosRepository) getAll() ([]Todo, error) {

	rows, err := r.db.Query("SELECT * FROM todos")

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

func (r *todosRepository) update(dto *UpdateTodoDto) (*Todo, error) {

	todo, err := r.getById(dto.Id)

	if err != nil {
		return nil, err
	}

	stmt, err := r.db.Prepare(
		"UPDATE todos SET name = ?, is_done = ? WHERE id = ?",
	)

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

func (r *todosRepository) delete(dto *DeleteTodoDto) (*Todo, error) {

	todo, err := r.getById(dto.Id)

	if err != nil {
		return nil, err
	}

	stmt, err := r.db.Prepare("DELETE FROM todos WHERE id = ?")

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
