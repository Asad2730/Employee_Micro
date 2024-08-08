package main

import (
	"log"

	"github.com/Asad2730/Employee_Micro/ApiGateway/controller"
	routes "github.com/Asad2730/Employee_Micro/ApiGateway/route"
	pb "github.com/Asad2730/Employee_Micro/Proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	grpcClinet := pb.NewEmployeeServiceClient(conn)
	empController := controller.NewEmployeeClient(grpcClinet)
	r := gin.Default()
	routes.RegisterUserRoutes(r, empController)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
