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
}
