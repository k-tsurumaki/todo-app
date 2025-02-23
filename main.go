package main

import (
	"fmt"

	"github.com/k-tsurumaki/todo-app/config"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SLQDriver)
	fmt.Println(config.Config.DBName)
	fmt.Println(config.Config.LogFile)
}