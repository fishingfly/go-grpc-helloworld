/*
 * @Author: zhounanjun
 * @Date: 2022-04-24 22:13:24
 * @LastEditors: zhounanjun
 * @LastEditTime: 2022-04-24 22:14:54
 * @Description: 请填写简介
 */
package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "helloworld/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world nanjun"
)

var (
	addr = flag.String("addr", "localhost:9001", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
