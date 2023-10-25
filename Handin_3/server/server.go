package main

import (
	"io"
	"log"
	"net"
	"sync"

	gRPC "github.com/MohaAmmar/DisysMMA/Handin_3/proto"

	"google.golang.org/grpc"
)

type Server struct {
	gRPC.UnimplementedChatServiceServer

	name  string
	port  string
	mutex sync.Mutex
}

func (s *Server) SayHi(msgStream gRPC.ChatService_SayHiServer) error {
	for {
		msg, err := msgStream.Recv()

		if err != io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("Received message from %s %s", msg.clientName, msg.message)
	}

	ack := &gRPC.Farewell{message: "Bye"}
	msgStream.SendAndClose(ack)

	return nil
}

func main() {
	list, _ := net.Listen("tcp", "localhost:5400")

	grpcServer := grpc.NewServer()

	server := Server{name: "Server", port: "5400"}

	gRPC.RegisterChatServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
