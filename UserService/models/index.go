package models

type User struct {
	Name     string `json:"name,omitempty" bson:"name"`
	Email    string `json:"email,omitempty" bson:"email" `
	Password string `json:"password,omitempty" bson:"pw"`
	Contact  int64  `json:"contact,omitempty" bson:"contact"`
	Role     []string `json:"role,omitempty" bson:"role"`
}


type List struct {
	Role           string `json:"role,omitempty" bson:"role"`
	Responsibility string `json:"responsibility,omitempty" bson:"responsibility"`
	Access         string `json:"access,omitempty" bson:"access"`
}

type UpdateRole struct {
	Name string `json:"name,omitempty" bson:"name"`
	Role []string `json:"role,omitempty" bson:"role"`
}

type AssociateRole struct {
	Name string `json:"name,omitempty" bson:"name"`
	Role string `json:"role,omitempty" bson:"role"`
}


