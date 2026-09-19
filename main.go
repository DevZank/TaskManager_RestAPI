package main

import (
	"net/http"

	"github.com/DevZank/TaskManagerAPI/config"
	"github.com/DevZank/TaskManagerAPI/handlers"
	"github.com/DevZank/TaskManagerAPI/models"
	"github.com/gorilla/mux"
)

func main() {
	dbConn := config.SetupDB()

	_, err := dbConn.Exec(models.CreateTableSQL)
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	taskHandler := handlers.NewTaskHandler(dbConn)

	router.HandleFunc("/tasks", taskHandler.ReadTasks).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	// DEFER: Quando a função main() [Função em que ele está em volta] parar de executar: chama a função Close para fechar a conexão com o banco de dados
	defer dbConn.Close()

	panic(http.ListenAndServe(":8080", router))
}
