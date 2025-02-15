package repository

import (
	"context"
	"errors"
	"log"
	"regexp"

	"github.com/GilangAndhika/elfume/config"
	"github.com/GilangAndhika/elfume/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword securely hashes a password
func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return ""
	}
	return string(hashedPassword)
}

// ComparePassword checks if a password matches the hashed version
func ComparePassword(hashedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

// EXISTING VALIDATION

// IsEmailExists checks if an email already exists in the database
func IsEmailExists(email string) (bool, error) {
	collection := config.MongoDB.Collection("users")

	var user model.User
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil // Email does not exist
		}
		return false, err // Database error
	}
	return true, nil // Email exists
}

// IsUsernameExists checks if a username already exists
func IsUsernameExists(username string) (bool, error) {
	collection := config.MongoDB.Collection("users")

	var user model.User
	err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil // Username does not exist
		}
		return false, err // Database error
	}
	return true, nil // Username exists
}

// IsRoleExists checks if a role exists
func IsRoleExists(name string) (bool, error) {
	collection := config.MongoDB.Collection("roles")

	var role model.Role
	err := collection.FindOne(context.TODO(), bson.M{"role_name": name}).Decode(&role)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil // Role does not exist
	} else if err != nil {
		return false, err // Database error
	}
	return true, nil // Role exists
}

// REGEXP VALIDATION

// IsEmailValid validates email format
func IsEmailValid(email string) bool {
	regex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	match, err := regexp.MatchString(regex, email)
	if err != nil {
		log.Println("Error matching email regex:", err)
		return false
	}
	return match
}

// IsPhoneValid validates Indonesian phone number format and converts it to international format
func IsPhoneValid(phone string) (bool, string) {
	regex := `^08[1-9][0-9]{7,13}$`
	match, err := regexp.MatchString(regex, phone)
	if err != nil {
		log.Println("Error matching phone regex:", err)
		return false, ""
	}
	if match {
		// Convert to international format (Indonesia)
		return true, "62" + phone[1:]
	}
	return false, ""
}
