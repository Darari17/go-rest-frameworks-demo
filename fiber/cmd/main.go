package main

import "github.com/go-rest-frameworks-demo/fiber/routes"

func main() {
	server := routes.NewServer()
	defer server.Close()

	server.Run(":3000")
}
