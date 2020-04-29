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
	Delay       = flag.Duration("delay", time.Second, "delay time Duration")
	Endpoints   = flag.String("endpoints", strings.Join(endPoints, ","), "etcd endpoints")
	Env         = flag.String("env", "", "dev in (local, dev) default local")
	ServiceName = flag.String("serviceName", "demand:engine/v1", "service name")
)

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

	r := etcd.NewResolver(useEndpoints)
	resolver.Register(r)
	// Use endpoint from "scheme://authority/endpoint" as the default
	//conn, err := grpc.Dial(r.Scheme()+"://authority/"+*ServiceName,
	//	grpc.WithBalancerName("round_robin"),
	//	//grpc.WithTimeout(time.Duration(time.Second*5)),
	//	grpc.WithInsecure())
	//
	//if err != nil {
	//	panic(err)
	//}
	//defer func() {
	//	if conn != nil {
	//		_ = conn.Close()
	//	}
	//}()

	pid := os.Getpid()

	ticker := time.NewTicker(*Delay)
	for t := range ticker.C {
		conn, err := grpc.Dial(r.Scheme()+"://authority/"+*ServiceName,
			grpc.WithBalancerName("round_robin"),
			//grpc.WithTimeout(time.Duration(time.Second*5)),
			grpc.WithInsecure())

		client := pb.NewGreeterClient(conn)
		resp, err := client.SayHello(context.Background(),
			&pb.HelloRequest{
				NodeName: fmt.Sprintf("client-go [pid=%v]", pid),
				Name:     "world " + strconv.Itoa(time.Now().Second())})
		if err == nil {
			fmt.Printf("%v: Reply is %s\n", t, resp.String())
		} else {
			fmt.Printf("call server error:%s\n", err)
		}
	}
}
