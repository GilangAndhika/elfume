package repository

import (
	"context"
	"elfume/config"
	"elfume/model"
	"log"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return ""
	}
	return string(hashedPassword)
}

func ComparePassword(hashedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}


// EMAIL AND USERNAME VALIDATION

// Check if email already exists
func IsEmailExists(email string) (bool, error) {
	client := config.GetDB()
	collection := client.Database("elfume").Collection("users")

	var user model.User
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false, nil // Email does not exist
		}
		return false, err // Database error
	}
	return true, nil // Email exists
}

// Check if username already exists
func IsUsernameExists(username string) (bool, error) {
	client := config.GetDB()
	collection := client.Database("elfume").Collection("users")

	var user model.User
	err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false, nil // Username does not exist
		}
		return false, err // Database error
	}
	return true, nil // Username exists
}


// REGEXP VALIDATION

// Email Regexp Validation
func IsEmailValid(email string) bool {
	regex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	match, _ := regexp.MatchString(regex, email)
	return match
}

// Phone Regexp Validation
func IsPhoneValid(phone string) (bool, string) {
	regex := `^08[1-9][0-9]{7,13}$`
	match, _ := regexp.MatchString(regex, phone)
	if match {
		// Convert to international format (Indonesia)
		return true, "62" + phone[1:]
	}
	return false, ""
}
