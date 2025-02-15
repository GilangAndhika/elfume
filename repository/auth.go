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
		"role_name":  user.RoleName,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})

	return err
}

// GetUserByEmail finds a user by email
func GetUserByEmail(email string) (*model.User, error) {
	collection := config.MongoDB.Collection("users")

	var user model.User
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // Return nil if no user is found
		}
		return nil, err // Return error if database issue
	}
	return &user, nil
}

// GetUserByUsername finds a user by username
func GetUserByUsername(username string) (*model.User, error) {
	collection := config.MongoDB.Collection("users")

	var user model.User
	err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // Return nil if no user is found
		}
		return nil, err // Return error if database issue
	}
	return &user, nil
}

// GetUserByID finds a user by ID
func GetUserByID(id string) (*model.User, error) {
	collection := config.MongoDB.Collection("users")

	// Convert string to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err // Return error if invalid ObjectID
	}

	// Find user by ID
	var user model.User
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // Return nil if no user is found
		}
		return nil, err // Return error if database issue
	}
	return &user, nil
}

// GetUserByEmailOrUsername finds a user by email or username
func GetUserByEmailOrUsername(email, username string) (*model.User, error) {
	collection := config.MongoDB.Collection("users")

	// Query to find user by email OR username
	filter := bson.M{
		"$or": []bson.M{
			{"email": email},
			{"username": username},
		},
	}

	var user model.User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // Return nil if no user is found
		}
		return nil, err // Return database error
	}
	return &user, nil
}
