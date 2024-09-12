package tests

import (
	"ComputerWorld_API/db/model"
	"ComputerWorld_API/tests/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestPostOrder(t *testing.T) {
	ts.ClearTable("orders")

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/order/",
	}

	cases := []helpers.TestCase{
		{
			TestName: "Can create an Order",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: model.Order{
				OrderRef:     "SGWTDF",
				OrderAmount:  3,
				ProductID:    2,
				ProductPrice: 700,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts:  []string{`"order_ref":"SGWTDF"`},
			},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			ts.ExecuteTest(t, &testCase)
		})
	}

}

func TestGetOrder(t *testing.T) {
	ts.ClearTable("orders")

	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/order",
	}

	order := &model.Order{
		OrderRef:     "3GNGKF",
		OrderAmount:  2,
		ProductID:    2,
		ProductPrice: 700,
	}
	ts.S.Database.Create(order)

	cases := []helpers.TestCase{
		{
			TestName: "Can retrieve order by id",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, order.OrderID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					order.OrderRef,
					fmt.Sprintf(`"order_id":%v`, order.OrderID),
				},
			},
		},
		{
			TestName: "Cannot retrieve order by id",
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

func TestPutOrder(t *testing.T) {
	ts.ClearTable("orders")

	request := helpers.Request{
		Method: http.MethodPut,
		Url:    "/order",
	}

	order := &model.Order{
		OrderRef:     "TESTREF",
		OrderAmount:  10,
		ProductID:    2,
		ProductPrice: 350,
	}
	ts.S.Database.Create(order)

	cases := []helpers.TestCase{
		{
			TestName: "Can update order by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, order.OrderID),
			},
			RequestBody: model.Order{
				OrderRef:     "VBJC53",
				OrderAmount:  5,
				ProductID:    1,
				ProductPrice: 700,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{fmt.Sprintf(`"order_ref":"VBJC53"`)},
			},
		},
		{
			TestName: "Cannot find and update order by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, 100000),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
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
				BodyParts: []string{
					order.OrderRef,
					fmt.Sprintf(`"order_id":%v`, order.OrderID),
				},
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
			},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			ts.ExecuteTest(t, &testCase)
		})
	}
}
