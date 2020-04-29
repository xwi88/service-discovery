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
	// local nodes
	endPoints = []string{
		"http://127.0.0.1:2379",
	}
	// dev cluster
	devEndPoints = []string{
		"http://10.14.41.51:2379",
		"http://10.14.41.52:2379",
		"http://10.14.41.53:2379",
	}

	Endpoints        = flag.String("endpoints", strings.Join(endPoints, ","), "etcd endpoints")
	Env              = flag.String("env", "", "dev in (local, dev) default local")
	ServiceName      = flag.String("serviceName", "demand:engine/v1", "service name")
	Port             = flag.Int("port", 50051, "listening port")
	NodeName         = flag.String("nodeName", "", "server name")
	RegisterInterval = flag.Duration("internal", time.Second*10, "register interval")

	replyAddr = ""
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
	pid int
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.String())
	return &pb.HelloReply{NodeName: fmt.Sprintf("server-go [pid=%v] %v", s.pid, *NodeName),
		Message: "Receive client msg: " + in.GetName(),
		Ip:      replyAddr}, nil
}

func main() {
	flag.Parse()

	useEndpoints := *Endpoints
	if Env != nil {
		if *Env == "dev" {
			useEndpoints = strings.Join(devEndPoints, ",")
		} else {
			useEndpoints = strings.Join(endPoints, ",")
		}
	}

	scheme := "http"
	localIP := utils.LocalIP()
	port := strconv.Itoa(*Port)
	serviceAddr := fmt.Sprintf("%v://%v:%v", scheme, localIP, port)
	replyAddr = serviceAddr

	serviceName := *ServiceName
	registerInterval := *RegisterInterval

	pid := os.Getpid()

	fmt.Printf("pid: %v, ip:%v, serviceAddr: %v, serviceName: %v, registerInterval: %v, endPoints:%v\n",
		pid, localIP, serviceAddr, serviceName, registerInterval, useEndpoints)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{pid: pid})

	go etcd.Register(useEndpoints, serviceName, serviceAddr, registerInterval)

	ch := make(chan os.Signal, 1)
	// SIGKILL 和 SIGSTOP Neither of these signals can be captured by the application,
	// nor can they be blocked or ignored by the operating system.
	// kill -9 pid => SIGKILL
	//os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for s := range ch {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
				etcd.UnRegister(serviceName, serviceAddr)
				os.Exit(int(s.(syscall.Signal)))
			case syscall.SIGUSR1:
				log.Printf("signal:usr1 %v", s)
			case syscall.SIGUSR2:
				log.Printf("signal:usr2 %v", s)
			default:
				log.Printf("signal:other %v", s)
			}
		}
	}()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
