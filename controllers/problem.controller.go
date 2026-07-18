package controllers

import (
	"coderX/middleware"
	"coderX/services"
	"coderX/models"
	"coderX/utils/fomatter"
	"net/http"
)

type ProblemController struct {
	service services.ProblemService
}

func NewProblemController (_service services.ProblemService) *ProblemController{
	return &ProblemController{
		service: _service,
	}
}

func (controller *ProblemController) CreateProblem(w http.ResponseWriter, r *http.Request) {
	
	payloadAny := r.Context().Value(middleware.PayloadContextKey)
	payload, ok := payloadAny.(models.Problem)
	
	if !ok {
		fomatter.ErrorResponse(w, http.StatusInternalServerError, "Invalid payload type in context", nil)
		return
	}

	response, err := controller.service.CreateProblem(r.Context(), &payload)

	if err != nil {
		fomatter.ErrorResponse(w, http.StatusInternalServerError, "Failed to create problem", err)
		return
	}

	fomatter.SucessResponse(w, http.StatusCreated, "Problem created successfully", response)
}