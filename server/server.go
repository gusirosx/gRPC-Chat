package main

import (
	"context"
	pb "gRPC-Guide/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterChatServiceServer(srv, &server{})
	log.Printf("server listening at %v", lis.Addr())
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}

func (s *server) SayHello(ctx context.Context, message *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from client: %s", message.Body)
	return &pb.Message{Body: "Hello From the Server!"}, nil
}
