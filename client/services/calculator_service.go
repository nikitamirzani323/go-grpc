package services

import (
	"context"
	"fmt"
)

type CalculatorServive interface {
	Hello(name string) error
}
type calculatorService struct {
	calculatorClient CalculatorClient
}

func NewCalculatorService(calculatorClient CalculatorClient) calculatorService {
	return calculatorService{calculatorClient}
}

func (base calculatorService) Hello(name, lname string) error {
	req := HelloRequest{
		Name:  name,
		Lname: lname,
	}
	res, err := base.calculatorClient.Hello(context.Background(), &req)
	if err != nil {
		return err
	}
	fmt.Printf("Service : Hello\n")
	fmt.Printf("Request : %v - %v\n", req.Name, req.Lname)
	fmt.Printf("Response : %v\n", res.Result)

	return nil
}
