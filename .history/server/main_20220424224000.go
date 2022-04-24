/*
 * @Author: zhounanjun
 * @Date: 2022-04-24 13:04:34
 * @LastEditors: zhounanjun
 * @LastEditTime: 2022-04-24 22:39:54
 * @Description: 请填写简介
 */

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "helloworld/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 9001, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) GetTestNum(ctx context.Context, in *pb.HelloRequest) (*pb.TestNumReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestNum not implemented")
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
