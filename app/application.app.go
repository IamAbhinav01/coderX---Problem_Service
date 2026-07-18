package app

import (
	"coderX/config/db"
	"coderX/config/env"
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
	server := http.Server{
		Addr:         addr,
		Handler:      http.DefaultServeMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server running on port ", server.Addr)

	return server.ListenAndServe()
}
