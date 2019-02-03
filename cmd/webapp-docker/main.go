package main

import "github.com/bakhtik/webapp/internal/app/server"

const configFile = "config/webapp-docker/config.json"

func main() {
	// config := &server.Config{
	// 	CacheAddr:  "redis:6379",
	// 	ServerAddr: ":8080",
	// }

	// read configuration
	var config server.Config
	config.FromFile(configFile)

	server := server.NewServer(&config)
	server.Start()
}
