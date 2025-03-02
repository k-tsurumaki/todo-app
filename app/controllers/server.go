package controllers

import (
	"net/http"

	"github.com/k-tsurumaki/todo-app/config"
)

func StartMainServer() error {
	http.HandleFunc("/", top) // urlの指定
	return http.ListenAndServe(":" + config.Config.Port, nil)
}