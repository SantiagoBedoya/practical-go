package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SantiagoBedoya/auth-ms/internal/model"
	"github.com/SantiagoBedoya/auth-ms/internal/pb"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	Store(context.Context, *model.User) error
	FindByEmail(context.Context, string) (*model.User, error)
}

type service struct {
	repo       Repository
	privateKey []byte
	publicKey  []byte
}

func NewService(repo Repository, privateKey, publicKey []byte) pb.AuthServer {
	return &service{repo, privateKey, publicKey}
}

func (s service) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	payload, err := validateToken(s.publicKey, req.AccessToken)
	if err != nil {
		log.Println(err)
		return &pb.ValidateResponse{StatusCode: http.StatusUnauthorized}, nil
	}
	return &pb.ValidateResponse{Email: fmt.Sprint(payload["sub"]), StatusCode: http.StatusOK}, nil
}

func (s service) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return &pb.AuthResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		}, nil
	}
	if err := s.repo.Store(ctx, &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hash),
	}); err != nil {
		log.Println(err)
		return &pb.AuthResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		}, nil
	}
	token, err := generateToken(s.privateKey, req.Email)
	if err != nil {
		return &pb.AuthResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		}, nil
	}
	return &pb.AuthResponse{AccessToken: token, StatusCode: http.StatusCreated}, nil
}

func (s service) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	user, err := s.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return &pb.AuthResponse{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		}, nil
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return &pb.AuthResponse{
			StatusCode: http.StatusUnauthorized,
		}, nil
	}
	token, err := generateToken(s.privateKey, user.Email)
	if err != nil {
		return &pb.AuthResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		}, nil
	}
	return &pb.AuthResponse{AccessToken: token, StatusCode: http.StatusOK}, nil
}

func validateToken(publicKey []byte, token string) (jwt.MapClaims, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, err
	}
	tok, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := tok.Claims.(jwt.MapClaims)
	fmt.Println(ok, tok.Valid)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

func generateToken(privateKey []byte, subject string) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	}
	now := time.Now().UTC()
	claims := jwt.StandardClaims{
		ExpiresAt: now.Add(24 * time.Hour).Unix(),
		Subject:   subject,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", err
	}
	return token, nil
}
