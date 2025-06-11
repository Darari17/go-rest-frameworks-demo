package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/config"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/routes"
)

func main() {
	cfg := config.Cfg
	server := routes.NewServer()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Shutting down gracefully...")
		server.Close()
		os.Exit(0)
	}()

	port := cfg.AppConfig.AppPort
	if port == "" {
		port = "3000"
	}
	log.Printf("🚀 %s running at http://localhost:%s\n", cfg.AppConfig.AppName, port)

	server.Run(":" + port)
}
