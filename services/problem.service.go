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
	GetProblem(ctx context.Context,problemID string)(*models.Problem,error) 
	GetAllProblems(ctx context.Context) ([] *models.Problem,error)
	UpdateProbelm(ctx context.Context, problemID string , problemPayload *models.Problem)(*models.Problem,error)
	DeleteProblem(ctx context.Context,problemID string) error
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
		return nil,err
	}
	return payload, nil
}

func (service *ProblemServiceImpl) GetProblem(ctx context.Context, problemID string) (*models.Problem,error){

	payload,err := service.problem_repo.GetProblem(ctx,problemID)
	if err != nil{
		log.Println("Failed to senf problemID from service -> repository")
		return nil,err
	}

	return payload,nil

}

func (service *ProblemServiceImpl) GetAllProblems(ctx context.Context) ([] *models.Problem,error){

	payload,err := service.problem_repo.GetAllProblems(ctx)
	if err!= nil{
		log.Println("Failed to fetch all problems")
		return nil,err
	}

	return payload,nil
}

func (service *ProblemServiceImpl) UpdateProbelm(ctx context.Context, problemID string , problemPayload *models.Problem)(*models.Problem,error){

	sanitizedDescription, err := utils.Sanitize_to_HTML(problemPayload.Description)
	
	if err != nil{
		log.Fatal("Failed to santise the markdown")
		return nil,err
	}

	problemPayload.Description = sanitizedDescription
	fmt.Println("Problem DATA IS \n ",problemPayload)

	payload,err := service.problem_repo.UpdateProbelm(ctx,problemID,problemPayload)
	if err!= nil{
		log.Println("Failed to update problem payload from service -> repository")
		return nil,err
	}
	return payload, nil
}

func (service *ProblemServiceImpl) DeleteProblem(ctx context.Context,problemID string) error{

	err:= service.problem_repo.DeleteProblem(ctx,problemID)
	if err != nil{
		log.Println("Error happend while deleting the data ",err)
		return err 
	}
	return nil;
}