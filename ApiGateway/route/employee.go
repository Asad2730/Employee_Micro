package routes

import (
	"github.com/Asad2730/Employee_Micro/ApiGateway/controller"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, empController *controller.EmployeeClient) {
	userRoutes := router.Group("/employee")
	{
		userRoutes.POST("/", empController.CreateEmployee)
		userRoutes.GET("/", empController.GetEmployees)
		userRoutes.GET("/:id", empController.CreateEmployee)
		userRoutes.PUT("/", empController.UpdateEmployee)
		userRoutes.DELETE("/:id", empController.DeleteEmployee)

	}
}
