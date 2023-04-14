package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type FunctionTestSuite struct {
	suite.Suite
}

func (suite *FunctionTestSuite) SetupSuite() {
}

func (suite *FunctionTestSuite) TestRetrive() {
	a := suite.Assert()
	//data := createOrder()
	//err := Create(context.Background(), data)
	err := Retreive(context.Background(), "policies:OR-012")

	a.Nil(err)
}

func TestFunctionUnit(t *testing.T) {
	suite.Run(t, new(FunctionTestSuite))
}

func createOrder() *Order {
	data := Order{
		ID:       23,
		RefCode:  "test2",
		PolicyID: 5,
		Status:   "REQUESTED",
	}

	return &data
}
