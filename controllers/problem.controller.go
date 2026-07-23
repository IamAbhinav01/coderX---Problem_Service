package controllers

import (
	dto "coderX/DTO"
	"coderX/middleware"
	"coderX/models"
	"coderX/services"
	"coderX/utils/fomatter"
	"net/http"

	"github.com/go-chi/chi/v5"
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

func (controller *ProblemController) GetProblem(w http.ResponseWriter, r *http.Request){
	
	problemID := chi.URLParam(r,"id")
	if problemID == ""{
		fomatter.ErrorResponse(w,http.StatusBadRequest,"ProblemID is required",nil)
		return
	}

	response,err := controller.service.GetProblem(r.Context(),problemID)
	if err != nil{
		fomatter.ErrorResponse(w,http.StatusBadGateway,"Problem not found",err)
		return
	}

	fomatter.SucessResponse(w,http.StatusAccepted,"Problem retrieved successfully",response)

}

func (controller *ProblemController) GetAllProblems(w http.ResponseWriter, r *http.Request){

	response,err := controller.service.GetAllProblems(r.Context())
	if err != nil{
		fomatter.ErrorResponse(w,http.StatusBadGateway,"Unable to fetch all the prblems",err)
		return
	}

	fomatter.SucessResponse(w,http.StatusAccepted,"Problems retrieved succesfully",response)
}

func (controller *ProblemController) UpdateProblem(w http.ResponseWriter, r *http.Request){

	problemID := chi.URLParam(r,"id")
	if problemID == ""{
		fomatter.ErrorResponse(w,http.StatusBadRequest,"ProblemID is required",nil)
		return
	}

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

	response, err := controller.service.UpdateProbelm(r.Context(), problemID,&modelPayload)

	if err != nil {
		fomatter.ErrorResponse(w, http.StatusInternalServerError, "Failed to update problem", err)
		return
	}

	fomatter.SucessResponse(w, http.StatusOK, "Problem updated successfully", response)
}