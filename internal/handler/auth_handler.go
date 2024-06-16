package handler

import (
	pb "auth-service/api"
	"auth-service/internal/service"
	"context"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	err := h.authService.SignUp(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.SignUpResponse{Message: "User created successfully"}, nil
}

func (h *AuthHandler) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	token, err := h.authService.SignIn(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.SignInResponse{Token: token}, nil
}
