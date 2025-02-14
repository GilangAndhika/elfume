package model

import "github.com/golang-jwt/jwt/v4"

type JWTClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
	RoleID string `json:"role_id"`
}
