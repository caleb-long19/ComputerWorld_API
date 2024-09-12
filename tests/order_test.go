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
		{
			TestName: "404 Error: Failed to retrieve order by id",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, 1000000),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Error: Order with ID was not found",
			},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			ts.ExecuteTest(t, &testCase)
		})
	}
}

func TestPostOrder(t *testing.T) {
	ts.ClearTable("orders")

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/order/",
	}

	cases := []helpers.TestCase{
		{
			TestName: "Test 1 - Creating an Order",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: &model.Order{
				OrderRef:     "SGWTDF",
				OrderAmount:  3,
				ProductID:    2,
				ProductPrice: 700,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyPart:   "Order created successfully",
			},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			ts.ExecuteTest(t, &testCase)
		})
	}

}

func TestPutOrder(t *testing.T) {
	ts.ClearTable("orders")

	request := helpers.Request{
		Method: http.MethodPut,
		Url:    "/order",
	}

	ord := &model.Order{
		OrderRef:     "TESTREF",
		OrderAmount:  10,
		ProductID:    2,
		ProductPrice: 350,
	}
	ts.S.Database.Create(ord)

	cases := []helpers.TestCase{
		{
			TestName: "Test 1 - Update Order by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, ord.OrderID),
			},
			RequestBody: model.Order{
				OrderRef:     "VBJC53",
				OrderAmount:  5,
				ProductID:    1,
				ProductPrice: 700,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Order updated successfully",
			},
		},
		{
			TestName: "Test 2 - Error 404: Fail to find and update order by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, 100000),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Error: Could not find order by that ID",
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
	ts.ClearTable("orders")

	request := helpers.Request{
		Method: http.MethodDelete,
		Url:    "/order",
	}

	order := &model.Order{
		OrderRef:     "TESTREF",
		OrderAmount:  15,
		ProductID:    1,
		ProductPrice: 1200,
	}
	ts.S.Database.Create(order)

	cases := []helpers.TestCase{
		{
			TestName: "Test 1 - Delete order by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, order.OrderID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Success: Order has been deleted",
			},
		},
		{
			TestName: "Test 2 - Error 404: Fail to find and delete order by ID",
			Request: helpers.Request{
				Method: http.MethodDelete,
				Url:    "/order/1000",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Error: Could not find order by that ID",
			},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			ts.ExecuteTest(t, &testCase)
		})
	}
}
