package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = os.Getenv("JWT_SECRET")

func GenerateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Function to extract the user ID from the JWT token
func ExtractUserIDFromToken(tokenString string) (string, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the token signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil // Return the secret key for verification
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token: %v", err)
	}

	// Extract the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("unable to parse claims")
	}

	// Extract the userID from claims directly as string
	userID, ok := claims["sub"].(string)
	if !ok {
		return "", fmt.Errorf("sub not found in claims or is not a string")
	}

	return userID, nil // Return the user ID as a string
}
