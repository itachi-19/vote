package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "vote/auth" // Update the import path accordingly
)

type server struct {
	pb.UnimplementedAuthServiceServer
}

func (s *server) Authenticate(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	if req.Username == "user" && req.Password == "password" {
		return &pb.AuthResponse{
			Token:   "some-auth-token",
			Message: "Authentication successful",
			Success: true,
		}, nil
	}
	return &pb.AuthResponse{
		Message: "Invalid credentials",
		Success: false,
	}, nil
}

func (s *server) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	if req.Token == "some-auth-token" {
		return &pb.ValidateTokenResponse{
			Valid:  true,
			UserId: "12345",
		}, nil
	}
	return &pb.ValidateTokenResponse{
		Valid: false,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &server{})
	log.Printf("Auth server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
