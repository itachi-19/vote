package auth

import (
	"context"
  "crypto/md5"
  "encoding/hex"
  "vote/dao"
)

type server struct {
	UnimplementedAuthServiceServer
}

func (s *server) Authenticate(ctx context.Context, req *AuthRequest) (*AuthResponse, error) {
	if isValidLogin(req.Username, req.Password) {
		return &AuthResponse{
			Token:   dao.GetAuthZToken(req.Username),
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

func HashPassword(passwd string) string {
  hash := md5.Sum([]byte(passwd))
  return hex.EncodeToString(hash[:])
}

func isValidLogin(username string, passwd string) bool {
  user, ok := dao.Users[username]
  if ok && user.PasswordHash == HashPassword(passwd) {
    return true
  }
  return false
}
