package handlers

import "net/http"

type TaskService interface {
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
func PutTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
