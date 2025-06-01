package main

import (
	"log"
	"net/http"

	"github.com/mcutulo82/todo-api/internal/config"
	"github.com/mcutulo82/todo-api/internal/handler"
	"github.commcutulo82/todo-api/internal/repository"
	"github.commcutulo82/todo-api/internal/service"
)

func main() {
	db := config.ConnectDB()

	repo := repository.NewTaskRepository(db)
	svc := service.NewTaskService(repo)
	h := handler.NewTaskHandler(svc)

	http.HandleFunc("/tasks", h.HandleTasks)

	log.Println("Servidor em execução na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
