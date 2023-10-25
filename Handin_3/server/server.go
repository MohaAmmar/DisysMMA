package main

import (
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
