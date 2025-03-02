package controllers

import (
	"html/template"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("app/views/templates/top.html") // 引数で渡したテンプレートファイルを解析
	t.Execute(w, nil)
}