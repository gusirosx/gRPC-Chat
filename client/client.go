package main

import (
	pb "gRPC-Guide/proto"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	//var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)

	response, err := client.SayHello(context.Background(), &pb.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
