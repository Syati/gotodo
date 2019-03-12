package main

import (
	"flag"
	"fmt"
	"github.com/Syati/gotodo/db"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	//	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	//	if err != nil {
	//		log.Fatalf("failed to listen: %v", err)
	//	}

	db := db.GetDb()
	defer db.Close()

	// enable tls ref https://github.com/grpc/grpc-go/blob/master/examples/route_guide/server/server.go
	fmt.Println(fmt.Sprintf("Start grpc server at localhost:%d", *port))
}
