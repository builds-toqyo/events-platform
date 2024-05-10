package server

import (
	"log"
)

type MyServer struct {
	Logger *log.Logger
}

func NewMyServer()MyServer{
	return MyServer{}
}