package tests

import (
	"ComputerWorld_API/db/models"
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
			RequestBody: models.Order{
				OrderRef:    "SGWTDF",
				OrderAmount: 3,
				ProductID:   2,
				OrderPrice:  700,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts:  []string{`"order_ref":"SGWTDF"`},
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
				OrderPrice:  700,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot create an order as no Order Amount was given",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Order{
				OrderRef:    "SGWTDF",
				OrderAmount: 0,
				ProductID:   2,
				OrderPrice:  700,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot create an order as no ProductID was given",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Order{
				OrderRef:    "SGWTDF",
				OrderAmount: 3,
				ProductID:   0,
				OrderPrice:  700,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			//TODO: The price should be setup automatically, this test will become obsolete, only using it for now!
			TestName: "Cannot create an order as no Order Price was given",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Order{
				OrderRef:    "SGWTDF",
				OrderAmount: 3,
				ProductID:   2,
				OrderPrice:  0,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
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

	order := &models.Order{
		OrderRef:    "3GNGKF",
		OrderAmount: 2,
		ProductID:   2,
		OrderPrice:  700,
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
			TestName: "Can retrieve order by ID as a string was given",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, "1"),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
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

	order := &models.Order{
		OrderRef:    "TESTREF",
		OrderAmount: 10,
		ProductID:   2,
		OrderPrice:  350,
	}
	ts.S.Database.Create(order)

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
				OrderPrice:  700,
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

	order := &models.Order{
		OrderRef:    "TESTREF",
		OrderAmount: 15,
		ProductID:   1,
		OrderPrice:  1200,
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
			TestName: "Cannot find and delete order by ID",
			Request: helpers.Request{
				Method: http.MethodDelete,
				Url:    "/order/1000",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			TestName: "Cannot delete order as no ID was given",
			Request: helpers.Request{
				Method: request.Method,
				Url:    "/order/",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusMethodNotAllowed,
			},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			ts.ExecuteTest(t, &testCase)
		})
	}
}
