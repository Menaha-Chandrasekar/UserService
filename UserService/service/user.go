package service

import (
	"context"
	"errors"
	"fmt"

	//"fmt"
	"log"
	"userservice/interfaces"
	"userservice/models"
	"userservice/validation"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

type UserService struct {
	UserCollection    *mongo.Collection
	FeatureCollection *mongo.Collection
	ctx               context.Context
}

func InitService(collection *mongo.Collection, fcollection *mongo.Collection, ctx context.Context) interfaces.Iuser {
	userCollection = collection
	return &UserService{collection, fcollection, ctx}
}

func  CheckIfRecordExists(user *models.User) (bool) {
	// Define your filter (query)
	filter := bson.M{"email": user.Email}

	// Perform the FindOne query
	 var result bson.M
	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		// Handle the case where the document doesn't exist
		fmt.Println("Document not found0")
		return false
	} else if err != nil {
		// Handle other errors
		fmt.Println("uma")
		log.Fatal(err)
		return false
	} else {
		// Document found, you can work with 'result'
		fmt.Println(result)
		return true
	}
}


func (u *UserService) Createuser(user *models.User) (string, error) {
	scrambledPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(scrambledPassword)

	//check for duplicates
	check0 := CheckIfRecordExists(user)

	if !check0{
		//no existing doc, insert new
		fmt.Println("Document not found1")
		check1 := validation.ValidName(user.Name)
		check2 := validation.ValidEmail(user.Email)
		check3 := validation.ValidContact(user.Contact)

		if check1 && check2 && check3 {
			_, err := u.UserCollection.InsertOne(u.ctx, &user)

			if err != nil {
				return "nil", err
			}
			return "successfully user created", nil

		}
		if !check1 {
			return "nil", errors.New("enter valid name")
		}
		if !check2 {
			return "nil", errors.New("enter valid email")
		} else {
			return "nil", errors.New("enter valid contact")
		}
	}else{
		//document exits, so don't insert
		fmt.Println("Document found1")
		err := errors.New("Document already exists")
		return "nil", err
	}
}

func (u *UserService) ListFeatures(feature *models.List) (*models.List, error) {
	filter := bson.M{"role": feature.Role}
	var list *models.List
	res := u.FeatureCollection.FindOne(u.ctx, filter)
	err := res.Decode(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (u *UserService) FindEnabledUser(status *models.User)(*models.User,error){
	filter:=bson.M{"status":status.Status}
	var list *models.User
	res := u.UserCollection.FindOne(u.ctx, filter)
	err := res.Decode(&list)
	if err != nil {
		return nil, err
	}
	return list, nil

}

func (u *UserService) UpdateRole(role *models.UpdateRole) (string, error) {
	filter := bson.M{"name": role.Name}
	var update *models.UpdateRole
	res := u.UserCollection.FindOne(u.ctx, filter)
	err := res.Decode(&update)
	if err != nil {
		return "nil", err
	}
	_, err = u.UserCollection.UpdateOne(u.ctx, filter, bson.M{"$set": bson.M{"role": role.Role}})
	if err != nil {
		return "nil", err
	}
	return "Role Updated Sucessfully", err
}

func (u *UserService) DisableUser(user *models.User) (string, error) {
	filter := bson.M{"name": user.Name}
	_, err := u.UserCollection.UpdateOne(u.ctx, filter, bson.M{"$set": bson.M{"status": "disabled"}})
	if err != nil {
		return "nil", err
	}
	return "User is Disabled Sucessfully", err
}

func (u *UserService) EnableUser(user *models.User) (string, error) {
	filter := bson.M{"name": user.Name}
	_, err := u.UserCollection.UpdateOne(u.ctx, filter, bson.M{"$set": bson.M{"status": "Enabled"}})
	if err != nil {
		return "nil", err
	}
	return "User is Enabled Sucessfully", err
}

func (u *UserService) AssociateRole(role *models.AssociateRole) (string, error) {
	filter := bson.M{"name": role.Name}
	_, err := u.UserCollection.UpdateOne(u.ctx, filter, bson.M{"$push": bson.M{"role": role.Role}})
	if err != nil {
		return "nil", err
	}
	return "Role Associated Sucessfully", err
}

func (u *UserService) DeleteUser( user *models.User)(string,error){

	filter := bson.M{"email":user.Email}
	_,err:=u.UserCollection.DeleteOne(u.ctx,filter)
	if err != nil {
		return "nil", err
	}
	return "Deleted Successfully", err

}