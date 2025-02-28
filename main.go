package main

import (
	"fmt"

	"github.com/k-tsurumaki/todo-app/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SLQDriver)
	// fmt.Println(config.Config.DBName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.DB)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "text@example.com"
	// u.Password = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// user, _ := models.GetUser(1)
	// user.CreateTodo("Hello World")

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// user, _ := models.GetUser(3)
	// user.CreateTodo("aaa")

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// todosByUser, _ := user.GetTodosByUser()
	// for _, v := range todosByUser {
	// 	fmt.Println(v)
	// }

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)
	// t.Content = "Update Todo"
	// t.UpdateTodo()

	u, _ := models.GetUser(1)
	u.CreateTodo("uhobubuojxeaa")
	// todos, _ := u.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	t, _ := models.GetTodo(8)
	fmt.Println(t)
	
	t.DeleteTodo()
	t2, _ := models.GetTodo(8)
	fmt.Println(t2)
}
