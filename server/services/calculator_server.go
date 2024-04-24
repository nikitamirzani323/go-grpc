package services

import (
	context "context"
	"fmt"
)

type calculatorServer struct {
}

func NewCalculatorServer() CalculatorServer {
	return calculatorServer{}
}
func (calculatorServer) mustEmbedUnimplementedCalculatorServer() {}
func (calculatorServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	result := fmt.Sprintf("Hello %v %v", req.Name, req.Lname)
	res := HelloResponse{
		Result: result,
	}
	return &res, nil
}
