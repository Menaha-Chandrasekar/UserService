package main

import (
	"context"
	"fmt"
	"net"
	"userservice/config"
	"userservice/constants"
	"userservice/controllers"
	"userservice/service"
	pb "userservice/proto"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initApp(client *mongo.Client){
	usercoll:=client.Database(constants.DataBaseName).Collection("users")
	rolecoll:=client.Database(constants.DataBaseName).Collection("features")
	ctx:=context.TODO()
	service:=service.InitService(usercoll,rolecoll,ctx)
	controllers.UserService =service
}

func main(){
	fmt.Println("Server Started..")
	
	//1.
	client,err:=config.ConectDatabase()
	if err!=nil{
		fmt.Println("Error connecting to DB")
		panic(err)
	}
	fmt.Println("DB connected")
	defer client.Disconnect(context.TODO())
	
	//2.
	initApp(client)
	
	//3.
	lis,err:=net.Listen("tcp", ":3000")
	if err!=nil{
		fmt.Println("Error Listening to Server")
		panic(err)
	}

	//4.
	s := grpc.NewServer()
	pb.RegisterUserserviceServer(s, &controllers.RPCServer{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}

}