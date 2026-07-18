package controllers

import (
	"coderX/middleware"
	"coderX/services"
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

func (controller *ProblemController) CreateProblem(w http.ResponseWriter,r *http.Request){
	
	payload := r.Context().Value(middleware.PayloadContextKey)
	response , err := controller.service.CreateProblem(r.Context(),payload)
}



func (c *ProblemController) CreateProblemHandler(w http.ResponseWriter, r *http.Request) {
	
	

	// 1. Parse JSON body
	if err := json.NewDecoder(r.Body).Decode(&problemPayload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// 2. Call service layer
	createdProblem, err := c.service.CreateProblem(r.Context(), &problemPayload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Return JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Success": true,
		"Message": "Problem created successfully",
		"data":    createdProblem,
	})
}