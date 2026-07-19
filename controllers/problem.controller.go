package controllers

import (
	dto "coderX/DTO"
	"coderX/middleware"
	"coderX/models"
	"coderX/services"
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
	
	dtoPayload := r.Context().Value(middleware.PayloadContextKey).(dto.GenerateProblem)
	
	var testCases []models.TestCase
	for _, tc := range dtoPayload.TestCases {
		testCases = append(testCases, models.TestCase{
			Input:  tc.Input,
			Output: tc.Output,
		})
	}

	var codeSnippets []models.CodeSnippet
	for _, cs := range dtoPayload.CodeSnippets {
		codeSnippets = append(codeSnippets, models.CodeSnippet{
			Language:     cs.Language,
			StartSnippet: cs.StartSnippet,
			MidSnippet:   cs.MidSnippet,
			EndSnippet:   cs.EndSnippet,
		})
	}

	// Create your final Model representation (Controller Layer)
	modelPayload := models.Problem{
		Title:        dtoPayload.Title,
		Description:  dtoPayload.Description,
		Difficulty:   dtoPayload.Difficulty,
		TestCases:    testCases, 
		CodeSnippets: codeSnippets, 
		Editorial:    dtoPayload.Editorial,
		Topic:        dtoPayload.Topic,
	}

	response, err := controller.service.CreateProblem(r.Context(), &modelPayload)

	if err != nil {
		fomatter.ErrorResponse(w, http.StatusInternalServerError, "Failed to create problem", err)
		return
	}

	fomatter.SucessResponse(w, http.StatusCreated, "Problem created successfully", response)
}