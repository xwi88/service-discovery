package main

import (
	pb "github.com/xwi88.com/service-discovery/pb/go"
)

func main() {
	request := pb.HelloRequest{}
	// request.OsType = int32(pb.OSType_IOS)
	request.OsType = 0
	// request.OsType = pb.OSType_IOS;
}
