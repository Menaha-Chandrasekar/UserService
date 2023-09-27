package config

import (
	"context"
	"log"
	"time"
	"userservice/constants"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConectDatabase()(*mongo.Client,error){
	ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
    mongoCollection:=options.Client().ApplyURI(constants.ConnectionString)
	mongoClient,err:=mongo.Connect(ctx,mongoCollection)
	if err!=nil{
		log.Fatal(err.Error())
		return nil,err
	}
	return mongoClient,err
	}    


	func GetCollection(client *mongo.Client,dbname string, collectionName string)(*mongo.Collection){
		collection:=client.Database(dbname).Collection(collectionName)
		return collection
	}