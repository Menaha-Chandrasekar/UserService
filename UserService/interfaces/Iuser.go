package interfaces

import "userservice/models"

type Iuser interface{
	// CheckIfRecordExists(user *models.User) (bool, error) 
	Createuser( user *models.User)(string,error)
	ListFeatures(Features *models.List)(*models.List,error)
	UpdateRole(role *models.UpdateRole)(string,error)
	DisableUser(user *models.User)(string,error)
	EnableUser(user *models.User)(string,error)
	AssociateRole(role *models.AssociateRole)(string,error)
	DeleteUser(user *models.User)(string,error)
	FindEnabledUser(user *models.User)(*models.User,error)
}