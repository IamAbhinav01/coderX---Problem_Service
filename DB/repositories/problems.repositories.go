package repositories

import (
	"coderX/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProblemRepository interface {
	CreateProblem() (*mongo.InsertOneResult, error)
	GetProblem()
	GetAllProblems()
	UpdateProbelm()
	DeleteProblem()
}

type ProblemRepositoryImpl struct {
	db *mongo.Collection
}

func(repo *ProblemRepositoryImpl) CreateProblem() (*mongo.InsertOneResult, error){

	var problem models.Problem
	result,err := repo.db.InsertOne(context.Background(),problem)
	if err != nil{
		log.Println("Error occured while storing the problem data to MongoDB")
		return nil,err
	}
	
	return result,nil

}