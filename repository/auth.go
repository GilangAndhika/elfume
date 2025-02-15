package repository

import (
	"context"
	"errors"
	"time"

	"github.com/GilangAndhika/elfume/config"
	"github.com/GilangAndhika/elfume/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateAccount handles user registration
func CreateAccount(user *model.User) error {
	// Get MongoDB connection
	usersCollection := config.MongoDB.Collection("users")
	rolesCollection := config.MongoDB.Collection("roles")

	// Assign a new ObjectID if not set
	if user.UserID.IsZero() {
		user.UserID = primitive.NewObjectID()
	}
	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	user.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	// Fetch role_name from roles collection using role_id
	var role model.Role
	err := rolesCollection.FindOne(context.TODO(), bson.M{"_id": user.RoleID}).Decode(&role)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("invalid role_id: role not found")
		}
		return err // Database error
	}
	user.RoleName = role.RoleName // Assign the fetched role_name

	// Insert user into the database
	_, err = usersCollection.InsertOne(context.TODO(), bson.M{
		"_id":        user.UserID,
		"username":   user.Username,
		"email":      user.Email,
		"password":   user.Password,
		"phone":      user.Phone,
		"role_id":    user.RoleID,
		"role_name":  user.RoleName, // Now automatically filled
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})

	return err
}

// // GetUserbyEmail finds a user by email
// func GetUserbyEmail(email string) (model.User, error) {
// 	client := config.GetDB()
// 	collection := client.Database("elfume").Collection("users")

// 	var user model.User
// 	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
// 	return user, err
// }

// // GetUserbyUsername finds a user by username
// func GetUserbyUsername(username string) (model.User, error) {
// 	client := config.GetDB()
// 	collection := client.Database("elfume").Collection("users")

// 	var user model.User
// 	err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
// 	return user, err
// }

// // GetUserbyID finds a user by ID
// func GetUserbyID(id string) (model.User, error) {
// 	client := config.GetDB()
// 	collection := client.Database("elfume").Collection("users")

// 	// Convert string to ObjectID
// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return model.User{}, err // Return an empty struct and an error
// 	}

// 	// Find user by ID
// 	var user model.User
// 	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
// 	return user, err
// }
