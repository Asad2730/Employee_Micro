package service

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Asad2730/Employee_Micro/Employee/data"
	pb "github.com/Asad2730/Employee_Micro/Proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedEmployeeServiceServer
	address string
}

func NewServer(address string) *server {
	return &server{address: address}
}

func (s *server) CreateEmployee(ctx context.Context, employee *pb.Employee) (*pb.Employee, error) {
	data.Emp_Arr = append(data.Emp_Arr, employee)
	return employee, nil
}

func (s *server) GetEmployees(ctx context.Context, empty *pb.EmptyRequest) (*pb.EmployeeList, error) {
	return &pb.EmployeeList{Employee: data.Emp_Arr}, nil
}

func (s *server) GetEmployee(ctx context.Context, EmpID *pb.EmployeeIdRequest) (*pb.Employee, error) {
	for _, emp := range data.Emp_Arr {
		if emp.Id == EmpID.Id {
			return emp, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "Employee not found with id: %d\n", EmpID.Id)
}

func (s *server) UpdateEmployee(ctx context.Context, updateEmployeeDto *pb.UpdateEmployeeRequest) (*pb.Employee, error) {
	for i, emp := range data.Emp_Arr {
		if emp.Id == updateEmployeeDto.Id {
			data.Emp_Arr[i] = &pb.Employee{
				Id:         emp.Id,
				Name:       updateEmployeeDto.Name,
				Email:      updateEmployeeDto.Email,
				Department: emp.Department,
			}
			return data.Emp_Arr[i], nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "Employee not found with id: %d\n", updateEmployeeDto.Id)
}

func (s *server) DeleteEmployee(ctx context.Context, EmpId *pb.EmployeeIdRequest) (*pb.DeleteResponse, error) {
	for i, emp := range data.Emp_Arr {
		if emp.Id == EmpId.Id {
			data.Emp_Arr = append(data.Emp_Arr[:i], data.Emp_Arr[i+1:]...)

			return &pb.DeleteResponse{
				Message: "Employee Reomved Successfully!",
			}, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "Employee not found with id: %d\n", EmpId.Id)
}

func (s *server) Start() error {
	listeners, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpc := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(grpc, &server{})
	fmt.Println("gRPC server is running at ", s.address)
	return grpc.Serve(listeners)

}
