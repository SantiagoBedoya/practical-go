package accounts

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// AuthServiceInstance is the implementation of AuthService interface
type AuthServiceInstance struct {
	Repository Repository
}

func getJwtToken(userID string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        userID,
		ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
	})
	return jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// SignIn is the implementation of AuthService's method
func (s AuthServiceInstance) SignIn(ctx context.Context, email string, password string) (string, error) {
	userID, userPassword, err := s.Repository.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password)); err != nil {
		return "", ErrInvalidCredentials
	}
	token, err := getJwtToken(userID)
	if err != nil {
		return "", err
	}
	return token, nil
}

// SignUp is the implementation of AuthService's method
func (s AuthServiceInstance) SignUp(ctx context.Context, email string, password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	userID, err := s.Repository.Create(ctx, email, string(hash))
	if err != nil {
		return "", err
	}

	token, err := getJwtToken(userID)
	if err != nil {
		return "", err
	}
	return token, nil
}
