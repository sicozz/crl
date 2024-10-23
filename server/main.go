package main

import (
	"context"
	"fmt"
	"log"
	"net"

	crl "github.com/sicozz/crl/api/v0"
	"google.golang.org/grpc"
)

type implGreeterServer struct {
	crl.GreeterServer
	prefix string
}

func (s *implGreeterServer) Hello(ctx context.Context, name *crl.Name) (*crl.Greeting, error) {
	return &crl.Greeting{Greeting: fmt.Sprintf("[%v] Hello %v. How are ya?", s.prefix, name.Name)}, nil
}

func main() {
	fmt.Println("Welcome to the Computer Racing League")
	lis, err := net.Listen("tcp", "[::]:50050")
	if err != nil {
		log.Fatalf(fmt.Sprintf("Cannot create listener: %v", err))
	}
	grpcServer := grpc.NewServer()
	crl.RegisterGreeterServer(grpcServer, &implGreeterServer{prefix: "CRL"})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(fmt.Sprintf("Server stoped or crashed: %v", err))
	}
}
