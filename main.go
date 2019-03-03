package main

import "github.com/Syati/gotodo/server"

func main() {
	r := server.Routes()
	r.Run(":8080")
}