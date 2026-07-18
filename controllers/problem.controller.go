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
	payload.Title = "Sample Demo Testing of Connection between cotroller,service,repository"
	payload.Description = "This is a **markdown** test description. Will it parse?"
	payload.Difficulty = "easy"
	payload.TestCases = []models.TestCase{
		{Input: "1 2", Output: "3"},
		{Input: "4 5", Output: "6"},
	}
	payload.CodeSnippets = []models.CodeSnippet{
		{Language: "go", StartSnippet: "func main() {", EndSnippet: "}"},
		{Language: "cpp", StartSnippet: "#include <bits/stdc++.h> using namespace std; int main() {", EndSnippet: "}"},
	}
	payload.Topic = "Testing the Connection"
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