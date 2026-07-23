package routers

import (
	"coderX/controllers"
	"coderX/middleware"

	"github.com/go-chi/chi/v5"
)

func ProblemRoutes(r chi.Router, pc *controllers.ProblemController){

	r.With(middleware.ProblemCreation).Post("/create-problem",pc.CreateProblem)
	r.Get("/{id}",pc.GetProblem)
	r.Get("/",pc.GetAllProblems)
	r.With(middleware.ProblemCreation).Put("/{id}",pc.UpdateProblem)
}