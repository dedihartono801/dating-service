package middleware

import (
	"errors"
	"net/http"
	"time"

	"dating-service/pkg/config"
	"dating-service/pkg/helper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Define the claims structure for the JWT
type Claims struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
	Gender     string `json:"gender"`
	IsVerified bool   `json:"is_verified"`
	IsPremium  bool   `json:"is_premium"`
}

// Define a function for generating a new JWT
func GenerateToken(id string, email string, gender string, isVerified, isPremium bool) (string, error) {
	expirationTime := time.Now().Add(time.Hour * time.Duration(24))
	claims := &Claims{
		Id:     id,
		Email:  email,
		Gender: gender,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
		IsVerified: isVerified,
		IsPremium:  isPremium,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(config.GetEnv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Define a middleware for verifying JWT authentication and expiration
func AuthUser(c *fiber.Ctx) error {
	// Get the Authorization header from the request
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		err := helper.Error(http.StatusUnauthorized, "Authorization header not found", errors.New("authorization header not found"))
		return helper.ResponseError(c, err)
	}

	// Verify that the Authorization header starts with "Bearer "
	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		err := helper.Error(http.StatusUnauthorized, "Invalid format authorization", errors.New("invalid format authorization"))
		return helper.ResponseError(c, err)
	}

	// Parse the JWT from the Authorization header
	tokenString := authHeader[len("Bearer "):]
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv("JWT_SECRET")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			err = helper.Error(http.StatusUnauthorized, "Invalid signature", err)
			return helper.ResponseError(c, err)
		}
		err = helper.Error(http.StatusUnauthorized, "Invalid token", err)
		return helper.ResponseError(c, err)
	}

	// Check if the JWT has expired
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		err = helper.Error(http.StatusUnauthorized, "Invalid token", errors.New("invalid token"))
		return helper.ResponseError(c, err)
	}
	if claims.ExpiresAt < time.Now().Unix() {
		err = helper.Error(http.StatusUnauthorized, "Expired token", errors.New("expired token"))
		return helper.ResponseError(c, err)
	}

	// Set the user ID in the context for future requests
	c.Locals("email", claims.Email)
	c.Locals("gender", claims.Gender)
	c.Locals("user-id", claims.Id)
	c.Locals("is-verified", claims.IsVerified)
	c.Locals("is-premium", claims.IsPremium)

	// Call the next middleware in the chain
	return c.Next()
}
