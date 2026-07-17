package repositories

import (
	"coderX/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProblemRepository interface {
	CreateProblem(ctx context.Context,problemPayload *models.Problem) (*models.Problem, error)
	// GetProblem()
	// GetAllProblems()
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
	problemPayload.ID = primitive.NewObjectID()

	_,err := repo.db.InsertOne(ctx,problemPayload)

	if err != nil{
		log.Println("Error occured while storing the problem data to MongoDB")
		return nil,err
	}

	return problemPayload,nil

}