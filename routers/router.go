package routers

import (
	"coderX/controllers"

	"github.com/go-chi/chi/v5"
)

func ProblemRouter(pc *controllers.ProblemController) *chi.Mux{
	router := chi.NewRouter()

	router.Route("/api/v1/problems",func(r chi.Router) {
		/* setup the problem routes*/
	})
	return  router
}