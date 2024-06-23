package voting_system

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "vote/auth"
  "vote/dao"
)

func Login(username string, password string) *dao.User {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewAuthServiceClient(conn)

	// Create a new context with a timeout.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the Authenticate method.
	req := &pb.AuthRequest{Username: username, Password: password}
	resp, err := client.Authenticate(ctx, req)
	if err != nil {
		log.Fatalf("could not authenticate: %v", err)
	}

	log.Printf("Response: %v %v ", resp.Message, resp.Success)
  if resp.Success == false {
    return nil
  }

  user := dao.GetUser(username)
  return &user
}
