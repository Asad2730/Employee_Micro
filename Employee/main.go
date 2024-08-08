package main

import (
	"log"

	"github.com/Asad2730/Employee_Micro/Employee/service"
)

func main() {
	grpcServer := service.NewServer(":8000")
	if err := grpcServer.Start(); err != nil {
		log.Fatalf("Failed to serve %v", err.Error())
	}
}
