package main

import (
	"context"
	pb "gRPC-Guide/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement ChatServer.
type server struct {
	// Embed the unimplemented server
	pb.UnimplementedChatServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("failed to listen on port 50050: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterChatServiceServer(srv, &server{})
	log.Printf("server listening at %v", lis.Addr())
	// Register reflection service on gRPC server.
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// SayHello implements chat.Server
func (s *server) SayHello(ctx context.Context, message *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from client: %s", message.Body)
	return &pb.Message{Body: "Hello From the Server!"}, nil
}
