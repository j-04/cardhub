package main

import "fmt"

type Config struct {
	*Server
}

type Server struct {
	Port int16 `yaml:"port"`
}

func (server *Server) GetHostAndPort() string {
	return fmt.Sprintf(":%d", server.Port)
}
