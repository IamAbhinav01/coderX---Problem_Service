package routers

import (
	"coderX/controllers"

	"github.com/go-chi/chi/v5"
)

func ProblemRoutes(r chi.Router, pc *controllers.ProblemController){

	r.Post("/create-problem",pc.CreateProblem)
}