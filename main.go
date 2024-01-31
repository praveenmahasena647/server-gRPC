package main

import (
	"context"
	"log"
	"net"

	grpc_play "github.com/praveenmahasena647/grpc_play/protos"
	"google.golang.org/grpc"
)

type GreetStruct struct {
	grpc_play.UnimplementedGreetServer
}

func (g *GreetStruct) SayHello(ctx context.Context, in *grpc_play.HelloReq) (*grpc_play.HelloRes, error) {
	var thing = &grpc_play.HelloRes{}
	thing.GreetStr = in.Name
	return thing, nil
}

func main() {
	var li, err = net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	var gs = grpc.NewServer()
	grpc_play.RegisterGreetServer(gs, &GreetStruct{})

	var status = make(chan error)

	go func(s chan error) {
		if err := gs.Serve(li); err != nil {
			status <- err
		}
	}(status)
	log.Fatalln(<-status)
}
