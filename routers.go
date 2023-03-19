package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

/**
 * Обработка маршрутов
 *
 */
func handleRouters() {
	router := mux.NewRouter()
	router.HandleFunc("/", viewTasks).Methods("GET")
	router.HandleFunc("/add", addTask).Methods("GET")
	router.HandleFunc("/view/{id:[0-9]+}", viewTask).Methods("GET")
	router.HandleFunc("/edit/{id:[0-9]+}", editTask).Methods("GET")

	router.HandleFunc("/api/post", postTask).Methods("POST")
	router.HandleFunc("/api/put/{id:[0-9]+}", putTask).Methods("POST")
	router.HandleFunc("/api/delete/{id:[0-9]+}", deleteTask).Methods("DELETE")

	http.Handle("/", router)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	http.ListenAndServe(":1617", nil)
}
