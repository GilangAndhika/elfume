package repository

import (
	"context"
	"errors"
	"fmt"
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

// GetUserByID retrieves a user from the database by their ID
func GetUserByID(id string) (*model.User, error) {
	// Convert string ID to primitive.ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID format: %v", err)
	}

	// Get database connection
	userCollection := config.MongoDB.Collection("users")

	// Find user by ID
	var user model.User
	err = userCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
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

// Get all users from the database
func GetAllUsers() ([]model.User, error) {
	// Get database connection
	userCollection := config.MongoDB.Collection("users")

	// Find all users
	cursor, err := userCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %v", err)
	}
	defer cursor.Close(context.Background())

	// Decode users
	var users []model.User
	if err = cursor.All(context.Background(), &users); err != nil {
		return nil, fmt.Errorf("failed to decode users: %v", err)
	}

	return users, nil
}

// UpdateUser updates an existing user's information by ID
func UpdateUser(id string, updatedUser model.User) error {
	// Convert ID to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid user ID format: %v", err)
	}

	// Get database connection
	userCollection := config.MongoDB.Collection("users")

	// Set updated timestamp
	updatedUser.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	// Define the update operation
	update := bson.M{
		"$set": bson.M{
			"username":  updatedUser.Username,
			"email":     updatedUser.Email,
			"phone":     updatedUser.Phone,
			"role_id":   updatedUser.RoleID,
			"updated_at": updatedUser.UpdatedAt,
		},
	}

	// Perform the update
	result, err := userCollection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	// Check if the user was found and modified
	if result.MatchedCount == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// DeleteUser deletes a user by ID
func DeleteUser(id string) error {
	// Convert string ID to primitive.ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid user ID format: %v", err)
	}

	// Get database connection
	userCollection := config.MongoDB.Collection("users")

	// Delete user by ID
	result, err := userCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	// Check if any document was actually deleted
	if result.DeletedCount == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}