package models

import (
	"time"

	"example.com/todo-app/db"
)

type Todo struct {
	Id          int64
	Title       string `binding:"required"`
	Description string `binding:"required"`
	DateTime    time.Time
	UserId      int64
}

func (todo *Todo) Save() error {
	query := `
	INSERT INTO todos(title, description, dateTime, user_Id)
	VALUES (?,?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(todo.Title, todo.Description, todo.DateTime, todo.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	todo.Id = id
	return err
}

func GetAllTodos() ([]Todo, error) {
	query := "SELECT * FROM todos"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.DateTime, &todo.UserId)

		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func GetTodoById(id int64) (*Todo, error) {
	query := "SELECT * FROM todos WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var todo Todo
	err := row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.DateTime, &todo.UserId)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (todo Todo) Update() error {
	query := `
	UPDATE todos
	SET title = ?, description = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Title, todo.Description, todo.DateTime, todo.Id)
	return err
}

func (todo Todo) Delete() error {
	query := "DELETE FROM todos WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Id)
	return err
}

func (todo Todo) LinkTodoToUser(userId int64) error {
	query := "INSERT INTO todoUsers(todo_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Id, userId)

	return err
}

func (todo Todo) UnlinkTodoFromUser(userId int64) error {
	query := "DELETE FROM todoUsers WHERE todo_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Id, userId)

	return err
}
