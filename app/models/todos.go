package models

import (
	"log"
	"time"
)

type Todo struct {
	ID int
	Content string
	UserID int
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (content, user_id, created_at) values (?, ?, ?)`

	_, err = DB.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTodo(id int) (todo Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos where id = ?`

	todo = Todo{}
	err = DB.QueryRow(cmd, id).Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)

	return todo, err
}

// func UpdateTodo()