package main

import "github.com/bakhtik/webapp/internal/app/server"

func main() {
	config := &server.Config{
		CacheAddr:  "redis:6379",
		ServerAddr: ":8080",
	}
	server := server.NewServer(config)
	server.Start()
}
