package controllers

import (
	"context"
	"userservice/interfaces"
	"userservice/models"
	u "userservice/proto"
)

//creating communication between proto and service

type RPCServer struct {
	u.UnimplementedUserserviceServer
}

var (
	UserService interfaces.Iuser
)

func (r *RPCServer) CreateUser(ctx context.Context, req *u.User) (*u.UserResponse, error) {
	//mapping the proto fields to the models
	dbuser := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Contact:  req.Contact,
		Role:     req.Role,
	}
	_, err := UserService.Createuser(dbuser)
	if err != nil {
		return nil, err
	} else {
		res := &u.UserResponse{
			Message: "success",
		}
		return res, nil
	}
}

func (r *RPCServer) ListFeatures(ctx context.Context, req *u.Feature) (*u.Featureresponse, error) {
	dblist := &models.List{
		Role: req.Role,
	}
	list, err := UserService.ListFeatures(dblist)
	if err != nil {
		return nil, err
	} else {
		res := &u.Featureresponse{
			Role:           list.Role,
			Responsibility: list.Responsibility,
			Access:         list.Access,
		}
		return res, nil
	}
}

func (r *RPCServer) UpdateRole(ctx context.Context, req *u.Role) (*u.UserResponse, error) {
	dbrole := &models.UpdateRole{
		Name: req.Name,
		Role: req.Role,
	}
	_, err := UserService.UpdateRole(dbrole)
	if err != nil {
		return nil, err
	} else {
		res := &u.UserResponse{
			Message: "success",
		}
		return res, nil
	}
}

func (r *RPCServer) DisableUser(ctx context.Context, req *u.UserRequest) (*u.UserResponse, error) {
	dbuser := &models.User{
		Name: req.Name,
	}
	_, err := UserService.DisableUser(dbuser)
	if err != nil {
		return nil, err
	} else {
		res := &u.UserResponse{
			Message: "success",
		}
		return res, nil
	}
}

func (r *RPCServer) EnableUser(ctx context.Context, req *u.UserRequest) (*u.UserResponse, error) {
	dbuser := &models.User{
		Name: req.Name,
	}
	_, err := UserService.EnableUser(dbuser)
	if err != nil {
		return nil, err
	} else {
		res := &u.UserResponse{
			Message: "success",
		}
		return res, nil
	}
}

func (r *RPCServer) AssociateRole(ctx context.Context, req *u.AssociateRequest) (*u.UserResponse, error) {
	dbuser := &models.AssociateRole{
		Name: req.Name,
		Role: req.Role,
	}
	_, err := UserService.AssociateRole(dbuser)
	if err != nil {
		return nil, err
	} else {
		res := &u.UserResponse{
			Message: "Associate success",
		}
		return res, nil
	}
}

func (r *RPCServer)DeleteUser(ctx context.Context, req *u.DeleteRequest)(*u.UserResponse,error){
   dbuser:=&models.User{
	Email: req.Email,
   }
   _,err:= UserService.DeleteUser(dbuser)
   if err != nil {
	return nil, err
} else {
	res := &u.UserResponse{
		Message: "Delete success",
	}
	return res, nil
}

}


func (r *RPCServer) FindEnabledUser(ctx context.Context, req *u.StatusRequest) (*u.User, error) {
	dbuser := &models.User{
		Status: req.Status,
	}
	list, err := UserService.FindEnabledUser(dbuser)
	if err != nil {
		return nil, err
	} else {
		res := &u.User{
			Name: list.Name,
			Email: list.Email,
			Password: list.Password,
			Contact: list.Contact,
			Role: list.Role,
			Status: list.Status,
		}
		
		return res, nil
	}
}


