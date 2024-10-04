package tests

import (
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/tests/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestOrderCreate(t *testing.T) {
	ts.ClearTable("orders")

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/order/",
	}

	ord := &models.Order{
		OrderRef:    "TESTREF",
		OrderAmount: 3,
		ProductID:   1,
	}
	ts.S.Database.Create(ord)

	cases := []helpers.TestCase{
		{
			TestName: "Can create an Order",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Order{
				OrderRef:    "SGWTDF",
				OrderAmount: 3,
				ProductID:   1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts:  []string{`"order_ref":"SGWTDF"`},
			},
		},
		{
			TestName: "Cannot create an order as order ref contains special characters",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Order{
				OrderRef:    "TESTREF##@",
				OrderAmount: 3,
				ProductID:   2,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot create an order as no Order Reference was given",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Order{
				OrderRef:    "",
				OrderAmount: 3,
				ProductID:   2,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot create an order as order ref already exists",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Order{
				OrderRef:    "TESTREF",
				OrderAmount: 3,
				ProductID:   2,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusConflict,
			},
		},
		{
			TestName: "Cannot create an order as Product ID does not exist",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Order{
				OrderRef:    "TestIDExists",
				OrderAmount: 3,
				ProductID:   99999999,
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

func TestOrderGet(t *testing.T) {
	ts.ClearTable("orders")

	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/order",
	}

	order := &models.Order{
		OrderRef:    "3GNGKF",
		OrderAmount: 2,
		ProductID:   2,
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
			TestName: "Cannot retrieve order that does not exist",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, 1000000),
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

func TestOrderUpdate(t *testing.T) {
	ts.ClearTable("orders")

	request := helpers.Request{
		Method: http.MethodPut,
		Url:    "/order",
	}

	order := &models.Order{
		OrderRef:    "TESTREF",
		OrderAmount: 10,
		ProductID:   1,
	}
	ts.S.Database.Create(order)

	orderDuplicate := &models.Order{
		OrderRef:    "TESTDUPE",
		OrderAmount: 5,
		ProductID:   1,
	}
	ts.S.Database.Create(orderDuplicate)

	cases := []helpers.TestCase{
		{
			TestName: "Can update order by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, order.OrderID),
			},
			RequestBody: models.Order{
				OrderRef:    "VBJC53",
				OrderAmount: 5,
				ProductID:   1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts:  []string{fmt.Sprintf(`"order_ref":"VBJC53"`)},
			},
		},
		{
			TestName: "Cannot find and update order by ID that does not exist",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, 100000),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			TestName: "Cannot update an order as order ref contains special characters",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, order.OrderID),
			},
			RequestBody: models.Order{
				OrderRef:    "TESTREF##@",
				OrderAmount: 3,
				ProductID:   1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot update an order as no Order Reference was given",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, order.OrderID),
			},
			RequestBody: models.Order{
				OrderRef:    "",
				OrderAmount: 3,
				ProductID:   1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot update an order as order ref already exists",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, orderDuplicate.OrderID),
			},
			RequestBody: models.Order{
				OrderRef:    "TESTDUPE",
				OrderAmount: 3,
				ProductID:   2,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusConflict,
			},
		},
		{
			TestName: "Cannot update an order as Product ID does not exist",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, order.OrderID),
			},
			RequestBody: models.Order{
				OrderRef:    "TestIDExists",
				OrderAmount: 3,
				ProductID:   99999999,
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

func TestOrderDelete(t *testing.T) {
	request := helpers.Request{
		Method: http.MethodDelete,
		Url:    "/order",
	}

	order := &models.Order{
		OrderRef:    "TESTREF",
		OrderAmount: 15,
		ProductID:   1,
	}
	ts.S.Database.Create(order)

	cases := []helpers.TestCase{
		{
			TestName: "Can delete order by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, order.OrderID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
			},
		},
		{
			TestName: "Cannot find and delete order that does not exist",
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
