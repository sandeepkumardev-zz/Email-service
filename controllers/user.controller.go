package controllers

import "net/http"

func HomeController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
