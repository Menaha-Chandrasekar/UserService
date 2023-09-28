package main

import (
	"context"
	"fmt"
	"net/http"
	pb "userservice/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client Started Running...")

	//1.
	r := gin.Default()
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//2.
	client := pb.NewUserserviceClient(conn)

	//3.
	r.POST("/add", func(c *gin.Context) {
		// var dbMutex sync.Mutex

		var request pb.User
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
		// Lock the mutex to ensure only one request can modify the database at a time.
		// dbMutex.Lock()
		// defer dbMutex.Unlock()
		

		//fup: internally client connects postman to controllers
		res, err := client.CreateUser(context.TODO(), &request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message})
	})

	r.POST("/update", func(c *gin.Context) {
		var request pb.Role
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		res, err := client.UpdateRole(context.TODO(), &request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message})
	})

	r.POST("/list", func(c *gin.Context) {
		var request pb.Feature
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		res, err := client.ListFeatures(context.TODO(), &request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Role": res.Role, "Responsibility": res.Responsibility, "Access": res.Access})
	})

	r.POST("/disable", func(c *gin.Context) {
		var request pb.UserRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		res, err := client.DisableUser(context.TODO(), &request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message})
	})

	r.POST("/enable", func(c *gin.Context) {
		var request pb.UserRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		res, err := client.EnableUser(context.TODO(), &request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message})
	})

	r.POST("/associaterole", func(c *gin.Context) {
		var request pb.AssociateRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		res, err := client.AssociateRole(context.TODO(), &request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message})
	})

	r.POST("/delete",func(c *gin.Context){
		var request pb.DeleteRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		res, err := client.DeleteUser(context.TODO(), &request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message})
	})

	r.POST("/findenableduser",func(c *gin.Context){
		var request pb.StatusRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		res, err := client.FindEnabledUser(context.TODO(), &request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Name": res.Name, "Email": res.Email, "Password": res.Password,"Contact":res.Contact,"Role":res.Role,"Status":res.Status})
	})
	r.Run(":4000")

}

