package config

import "fmt"

type Config struct {
	*Server
	*Database
}

type Server struct {
	Port int16 `yaml:"port"`
}

func (server *Server) GetHostAndPort() string {
	return fmt.Sprintf(":%d", server.Port)
}

type Database struct {
	Stub bool `yaml:"enable-stub"`
	*Cassandra
}

type Cassandra struct {
	Username          string `yaml:"username"`
	Password          string `yaml:"password"`
	Host              string `yaml:"host"`
	ConnectionTimeout int64  `yaml:"timeout"`
	Port              int16  `yaml:"port"`
}
