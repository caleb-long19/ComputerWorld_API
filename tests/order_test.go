package tests

import (
	"ComputerWorld_API/database/model"
	"ComputerWorld_API/tests/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestGetOrder(t *testing.T) {
	ts.ClearTable("orders")

	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/order",
	}

	ord := &model.Order{
		OrderRef:     "3GNGKF",
		OrderAmount:  2,
		ProductID:    2,
		ProductPrice: 700,
	}
	ts.S.Database.Create(ord)

	cases := []helpers.TestCase{
		// Cannot get manufacturer with invalid ID
		// Can get manufacturer
		{
			TestName: "Retrieve order by id",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, ord.OrderID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf("order_id %v", ord.OrderID),
				},
			},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			ts.ExecuteTest(t, &testCase)
		})
	}
}

func TestDeleteOrder(t *testing.T) {

	request := helpers.Request{
		Method: http.MethodDelete,
		Url:    "/order/1",
	}

	cases := []helpers.TestCase{
		{
			TestName: "Delete order by id",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Order has been deleted",
			},
		},
		//{
		//	TestName: "Return 404 if a string is used to delete a manufacturer",
		//	Request: helpers.Request{
		//		Method: request.Method,
		//		Url:    "/manufacturer/NAME",
		//	},
		//	Expected: helpers.ExpectedResponse{
		//		StatusCode: http.StatusNotFound,
		//	},
		//},
	}
	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			ts.ExecuteTest(t, &testCase)
		})
	}
}
