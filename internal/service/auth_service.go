package service

import (
	"auth-service/internal/model"
	"auth-service/internal/repository"
	"auth-service/pkg/jwt"
	"auth-service/pkg/kafka"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepository}
}

func (s *AuthService) SignUp(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		Username: username,
		Password: string(hashedPassword),
	}

	err = s.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	kafka.ProduceMessage("user_signup", username)
	return nil
}

func (s *AuthService) SignIn(username, password string) (string, error) {
	user, err := s.userRepository.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := jwt.GenerateJWT(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
