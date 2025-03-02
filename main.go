package main

import (
	"fmt"

	"github.com/k-tsurumaki/todo-app/app/controllers"
	"github.com/k-tsurumaki/todo-app/app/models"
)

func main() {
	fmt.Println(models.DB)
	controllers.StartMainServer()
}
