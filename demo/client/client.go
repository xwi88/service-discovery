// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/xwi88.com/service-discovery/discovery/etcd"
	"github.com/xwi88.com/service-discovery/pb/go/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

var (
	endPoints = []string{
		"http://127.0.0.1:2379",
	}
	Delay       = flag.Duration("delay", time.Second, "delay time Duration")
	Endpoints   = flag.String("endpoints", strings.Join(endPoints, ","), "etcd endpoints")
	ServiceName = flag.String("serviceName", "demand:engine", "service name")
)

func main() {
	flag.Parse()

	r := etcd.NewResolver(strings.Join(endPoints, ","))
	resolver.Register(r)
	// "://author/" ???
	conn, err := grpc.Dial(r.Scheme()+"://author/"+*ServiceName,
		grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	pid := os.Getpid()

	ticker := time.NewTicker(*Delay)
	for t := range ticker.C {
		client := pb.NewGreeterClient(conn)
		resp, err := client.SayHello(context.Background(),
			&pb.HelloRequest{NodeName: fmt.Sprintf("client-go [pid=%v]", pid),
				Name: "world " + strconv.Itoa(time.Now().Second())})
		if err == nil {
			fmt.Printf("%v: Reply is %s\n", t, resp.String())
		} else {
			fmt.Printf("call server error:%s\n", err)
		}
	}
}
