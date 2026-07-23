package repositories

import (
	"coderX/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProblemRepository interface {
	CreateProblem(ctx context.Context,problemPayload *models.Problem) (*models.Problem, error)
	GetProblem(ctx context.Context,problemID string)(*models.Problem,error) 
	GetAllProblems(ctx context.Context) ([] *models.Problem,error)
	// UpdateProbelm()
	// DeleteProblem()
}

type ProblemRepositoryImpl struct {
	db *mongo.Collection
}

func NewProblemsRepository(_db *mongo.Collection) *ProblemRepositoryImpl{
	return &ProblemRepositoryImpl{
		db: _db,
	}
}

func(repo *ProblemRepositoryImpl) CreateProblem(ctx context.Context,problemPayload *models.Problem) (*models.Problem, error){

	problemPayload.CreatedAt = time.Now()
	problemPayload.ID = primitive.NewObjectID().Hex() // this is actually make id a string

	_,err := repo.db.InsertOne(ctx,problemPayload)

	if err != nil{
		log.Println("Error occured while storing the problem data to MongoDB")
		return nil,err
	}

	return problemPayload,nil

}

func(repo *ProblemRepositoryImpl) GetProblem(ctx context.Context,problemID string) (*models.Problem,error){
	
	var problem *models.Problem

	err := repo.db.FindOne(ctx,bson.M{"_id":problemID}).Decode(&problem)
	if err != nil{
		log.Println("Error occured while retriving the user from the problem ID")
		return nil,err
	}

	return problem,nil;
}

func(repo *ProblemRepositoryImpl) GetAllProblems(ctx context.Context) ([] *models.Problem,error){
	
}