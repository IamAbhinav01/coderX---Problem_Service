package services

import "coderX/DB/repositories"

type ProblemService interface {
	CreateProblem()
}

type ProblemServiceImpl struct {
	problem_repo repositories.ProblemRepository
}

func NewProblemService(_problem_repo repositories.ProblemRepository) *ProblemServiceImpl{
	return &ProblemServiceImpl{
		problem_repo: _problem_repo,
	}
}

