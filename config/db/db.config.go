package db

import (
	"coderX/config/env"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInit() error{

	MONGO_DB := env.GetString("ATLAS_DB_URL")
	
	if MONGO_DB == ""{
		log.Fatal("Error while fetching MONGODB URI from env")
	}
	
	clientOption := options.Client().ApplyURI(MONGO_DB)
	client,err := mongo.Connect(context.Background(),clientOption)

	if err != nil{
		log.Fatal("Error while connecting to mongoDB",err)
	}

	err = client.Ping(context.Background(),nil)
	if err!= nil{
		log.Fatal("Unable to connect to MongoDB",err)
	}

	DB := client.Database("")
	
}

