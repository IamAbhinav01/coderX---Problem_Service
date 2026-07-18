package app

import (
	"coderX/DB/repositories"
	"coderX/config/db"
	"coderX/config/env"
	"coderX/controllers"
	"coderX/routers"
	"coderX/services"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type Application struct {
	Config string
}

func NewApplication() *Application {
	return &Application{
		Config: env.GetString("PORT"),
	}
}

func (app *Application) Run() error {

	err := db.DBInit()

	if err != nil {
		log.Fatal("Unable to Initialise the Database")
		return fmt.Errorf("Unable to connect to DataBase")
	}

	log.Printf("Database Initialised Succesully")

	addr := app.Config

	if strings.HasPrefix(addr, "") {
		addr = ":" + addr
	}
	problem_collection := db.Problems
	problem_repository := repositories.NewProblemsRepository(problem_collection)
	problem_service := services.NewProblemService(problem_repository)
	problem_controller := controllers.NewProblemController(problem_service)

	appRouter := routers.ProblemRouter(problem_controller)

	fmt.Println("configured application level stats")
	
	server := http.Server{
		Addr:         addr,
		Handler:      appRouter,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server running on port ", server.Addr)

	return server.ListenAndServe()
}
