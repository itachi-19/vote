package auth

import (
	"context"
)

type server struct {
	UnimplementedAuthServiceServer
}

func (s *server) Authenticate(ctx context.Context, req *AuthRequest) (*AuthResponse, error) {
	if req.Username == "user" && req.Password == "password" {
		return &AuthResponse{
			Token:   "some-auth-token",
			Message: "Authentication successful",
			Success: true,
		}, nil
	}
	return &AuthResponse{
		Message: "Invalid credentials",
		Success: false,
	}, nil
}

func (s *server) ValidateToken(ctx context.Context, req *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	if req.Token == "some-auth-token" {
		return &ValidateTokenResponse{
			Valid:  true,
			UserId: "12345",
		}, nil
	}
	return &ValidateTokenResponse{
		Valid: false,
	}, nil
}

func NewServer() *server {
  return &server{}
}
