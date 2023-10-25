package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	gRPC "github.com/MohaAmmar/DisysMMA/Handin_3/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var clientsName = flag.String("name", "default", "The name of the sender")
var serverPort = flag.String("port", "5400", "The port of the server")

var server gRPC.ChatServiceClient
var serverConn *grpc.ClientConn

func main() {
	//parse flag/arguments
	flag.Parse()

	fmt.Println("--- CLIENT APP ---")

	//log to file instead of console
	//f := setLog()
	//defer f.Close()

	//connect to server and close the connection when program closes
	fmt.Println("--- join Server ---")
	ConnectToServer()
	//defer ServerConn.Close()

	//start the biding
	//parseInput()
}

func ConnectToServer() {

	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:5400", opts...)

	if err != nil {
		log.Printf("Fail to Dial : %v", err)
		return
	}

	server = gRPC.NewChatServiceClient(conn)
	serverConn = conn
}

func SayHi() {
	stream, err := server.SayHi(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	stream.Send(&gRPC.Greeding{ClientName: *clientsName, Message: "Hi"})
	stream.Send(&gRPC.Greeding{ClientName: *clientsName, Message: "How are you?"})
	stream.Send(&gRPC.Greeding{ClientName: *clientsName, Message: "I'm fine, thanks."})

	farewell, err := stream.CloseAndRecv()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("server says: ", farewell)

}
