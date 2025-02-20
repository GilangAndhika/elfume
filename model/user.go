package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserID    primitive.ObjectID `json:"user_id" bson:"_id"`
	Username  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	Phone     string             `json:"phone" bson:"phone"`
	RoleID    primitive.ObjectID `json:"role_id" bson:"role_id"`
	RoleName  string             `json:"role_name" bson:"role_name"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

type Role struct {
	RoleID   primitive.ObjectID `json:"role_id" bson:"_id"`
	RoleName string             `json:"role_name" bson:"role_name"`
}

const (
	RoleAdmin    = "67aff19a533432bc3af88fe2"
	RoleCustomer = "67aff183533432bc3af88fe1"
)
