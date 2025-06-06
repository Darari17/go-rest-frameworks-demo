package main

import "github.com/Darari17/go-rest-frameworks-demo/echo/routes"

func main() {
	server := routes.NewServer()
	defer server.Close()

	server.Run(":3000")
}
