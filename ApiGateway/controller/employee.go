package controller

import (
	"strconv"

	pb "github.com/Asad2730/Employee_Micro/Proto"
	"github.com/gin-gonic/gin"
)

type EmployeeClient struct {
	grpcClient pb.EmployeeServiceClient
}

func NewEmployeeClient(client pb.EmployeeServiceClient) *EmployeeClient {
	return &EmployeeClient{
		grpcClient: client,
	}
}

func (client *EmployeeClient) CreateEmployee(c *gin.Context) {
	var emp pb.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(500, gin.H{"Error binding data": err.Error()})
		return
	}

	res, err := client.grpcClient.CreateEmployee(c, &emp)
	if err != nil {
		c.JSON(402, gin.H{"Error Creating": err.Error()})
		return
	}

	c.JSON(201, res)
}

func (client *EmployeeClient) GetEmployees(c *gin.Context) {
	res, err := client.grpcClient.GetEmployees(c, nil)
	if err != nil {
		c.JSON(500, gin.H{"Error Fetching": err.Error()})
		return
	}

	c.JSON(200, res.Employee)
}

func (client *EmployeeClient) GetEmployee(c *gin.Context) {
	id := c.Param("id")
	emp_id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(501, gin.H{"Error converting id": err})
		return
	}
	res, err := client.grpcClient.GetEmployee(c, &pb.EmployeeIdRequest{Id: int32(emp_id)})
	if err != nil {
		c.JSON(500, gin.H{"Error Fetching": err.Error()})
		return
	}

	c.JSON(200, res)
}

func (client *EmployeeClient) UpdateEmployee(c *gin.Context) {
	var emp pb.UpdateEmployeeRequest
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(500, gin.H{"Error binding data": err.Error()})
		return
	}
	res, err := client.grpcClient.UpdateEmployee(c, &emp)
	if err != nil {
		c.JSON(500, gin.H{"Error updating": err.Error()})
		return
	}

	c.JSON(200, res)
}

func (client *EmployeeClient) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	emp_id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(501, gin.H{"Error converting id": err})
		return
	}
	res, err := client.grpcClient.DeleteEmployee(c, &pb.EmployeeIdRequest{Id: int32(emp_id)})
	if err != nil {
		c.JSON(500, gin.H{"Error Deleting": err.Error()})
		return
	}

	c.JSON(200, res)
}
