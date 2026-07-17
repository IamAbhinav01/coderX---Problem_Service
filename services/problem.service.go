package services

import (
	"coderX/DB/repositories"
	"coderX/models"
	"coderX/utils"
	"context"
	"fmt"
	"log"
)

type ProblemService interface {
	CreateProblem(ctx context.Context,problemPayload *models.Problem) (*models.Problem, error)
}

type ProblemServiceImpl struct {
	problem_repo repositories.ProblemRepository
}

func NewProblemService(_problem_repo repositories.ProblemRepository) *ProblemServiceImpl{
	return &ProblemServiceImpl{
		problem_repo: _problem_repo,
	}
}

func (service *ProblemServiceImpl) CreateProblem(ctx context.Context,problemPayload *models.Problem) (*models.Problem, error){

	sanitizedDescription, err := utils.Sanitize_to_HTML(problemPayload.Description)
	
	if err != nil{
		log.Fatal("Failed to santise the markdown")
		return nil,err
	}

	problemPayload.Description = sanitizedDescription
	fmt.Println("Problem DATA IS \n ",problemPayload)

	payload,err := service.problem_repo.CreateProblem(ctx,problemPayload)
	if err!= nil{
		log.Println("Failed to send problem payload from service -> repository")
	}
	return payload, nil
}