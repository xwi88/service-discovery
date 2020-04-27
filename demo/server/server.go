// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/xwi88.com/service-discovery/discovery/etcd"
	"github.com/xwi88.com/service-discovery/pb/go/service"
	"github.com/xwi88.com/service-discovery/utils"
	"google.golang.org/grpc"
)

var (
	//endPoints = []string{
	//	"http://10.14.41.51:2379",
	//	"http://10.14.41.52:2379",
	//	"http://10.14.41.53:2379"}
	endPoints = []string{
		"http://127.0.0.1:2379",
	}
	Endpoints = flag.String("endpoints", strings.Join(endPoints, ","), "etcd endpoints")

	ServiceName      = flag.String("serviceName", "demand:engine", "service name")
	Port             = flag.Int("port", 50051, "listening port")
	NodeName         = flag.String("nodeName", "server", "server name")
	RegisterInterval = flag.Duration("internal", time.Second*10, "register interval")

	replyAddr = ""
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	log.Printf("Received: %v", in.String())
	return &pb.HelloReply{NodeName: *NodeName, Message: "Hello " + in.GetName(),
		Ip: replyAddr}, nil
}

func main() {
	flag.Parse()
	localIP := utils.LocalIP()
	port := strconv.Itoa(*Port)
	serviceAddr := fmt.Sprintf("%v:%v", localIP, port)
	replyAddr = serviceAddr

	serviceName := *ServiceName
	registerInterval := *RegisterInterval
	endPoints := *Endpoints

	fmt.Printf("ip:%v, serviceAddr: %v, serviceName: %v, registerInterval: %v,\nendPoints:%v\n",
		localIP, serviceAddr, serviceName, registerInterval, endPoints)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	go etcd.Register(endPoints,
		serviceName,
		serviceAddr,
		registerInterval,
	)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		etcd.UnRegister(serviceName, serviceAddr)

		if i, ok := s.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}

	}()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
