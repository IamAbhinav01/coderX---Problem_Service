package db

import (
	"coderX/config/env"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	Database *mongo.Database
	Problems *mongo.Collection
	Submissions *mongo.Collection
)


func DBInit() error{

	MONGO_DB := env.GetString("MONGO_DB_URL")
	db_name := env.GetString("DB_NAME")
	problem_collection := env.GetString("ProblemCollection")
	submission_collection := env.GetString("SubmissionCollection")
	
	if MONGO_DB == ""{
		log.Fatal("Error while fetching MONGODB URI from env")
		return fmt.Errorf("Error while fetching the mongoDB URI from env")
	}
	
	clientOption := options.Client().ApplyURI(MONGO_DB)
	client,err := mongo.Connect(context.Background(),clientOption)

	if err != nil{
		log.Fatal("Error while connecting to mongoDB",err)
		return fmt.Errorf("Error happend while connecting to MongoDB")
	}

	Client = client

	err = client.Ping(context.Background(),nil)
	if err!= nil{
		log.Fatal("Unable to connect to MongoDB",err)
		return fmt.Errorf("Unable to connect to MONGODB")
	}

	Database = client.Database(db_name)

	Problems = Database.Collection(problem_collection)
	Submissions = Database.Collection(submission_collection)

	return nil
}

