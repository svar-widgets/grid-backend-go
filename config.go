package main

import "svar-widgets/grid-backend-go/data"

type ConfigServer struct {
	URL  string
	Port string
	Cors []string
}

type AppConfig struct {
	Server     ConfigServer
	DB         data.DBConfig
	BinaryData string
}
