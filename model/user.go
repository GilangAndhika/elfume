package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"user_id" bson:"_id"`
	Username  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	Phone     string             `json:"phone" bson:"phone"`
	RoleID    string             `json:"role_id" bson:"role_id"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

type Role struct {
	ID       string `json:"role_id" bson:"_id"`
	RoleName string `json:"role_name" bson:"role_name"`
}

const (
	RoleAdmin    = "1"
	RoleCustomer = "2"
)
