// goroutine project main.go
package main

import (
	"fmt"
	"grpc"
	"http"
	"net"
	"pb"
	"runtime"
)

func main() {
	fmt.Println("Hello World!")
	runtime.SetBlockProfileRate(1)
	go func() {
		http.ListenAndServe(":10001", nil)
	}()

	lis, err := net.Listen("tcp", 10000)
	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, &routeGuideServer{})
	grpcServer.Serve(lis)
}
