package controllers

import (
	"log"
	"net/http"
)

func (m *Manager) GET_HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
	err := m.UserRepo.CreateUser("adilhan", "bigxxby@yandex.ru")
	if err != nil {
		log.Println(err.Error())
		return
	} else {
		log.Println("200")
	}
}
