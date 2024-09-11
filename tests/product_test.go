package tests

import (
	"ComputerWorld_API/database/model"
	"ComputerWorld_API/tests/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestGetProduct(t *testing.T) {
	ts.ClearTable("products")

	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/product",
	}

	pd := &model.Product{
		ProductName:    "Xbox Series Y",
		ProductCode:    "CHZXMGJ",
		ManufacturerID: 1,
		Stock:          250,
		Price:          400,
	}
	ts.S.Database.Create(pd)

	cases := []helpers.TestCase{
		// Cannot get manufacturer with invalid ID
		// Can get manufacturer
		{
			TestName: "Retrieve product by id",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, pd.ProductID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf("product_id %v", pd.ProductID),
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

//func TestPostProduct(t *testing.T) {
//	ts.ClearTable("products")
//
//	request := helpers.Request{
//		Method: http.MethodPost,
//		Url:    "/product",
//	}
//
//	cases := []helpers.TestCase{
//		{
//			TestName: "Test 1 - Create Product",
//			Request: helpers.Request{
//				Method: request.Method,
//				Url:    request.Url,
//			},
//			Expected: helpers.ExpectedResponse{
//				StatusCode: http.StatusCreated,
//				BodyPart:   "Product created successfully",
//			},
//		},
//	}
//	for _, testCase := range cases {
//		t.Run(testCase.TestName, func(t *testing.T) {
//			ts.ExecuteTest(t, &testCase)
//		})
//	}
//
//}

func TestPutProduct(t *testing.T) {
	ts.ClearTable("products")

	request := helpers.Request{
		Method: http.MethodPut,
		Url:    "/product/1",
	}

	cases := []helpers.TestCase{
		{
			TestName: "Test 1 - Update Product by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: model.Product{
				ProductCode:    "CHZXMG45J",
				ProductName:    "Xbox 1080",
				ManufacturerID: 1,
				Stock:          450,
				Price:          250,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Successfully updated product",
			},
		},
		{
			TestName: "Test 2 - Error 404: Fail to update product by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, 100000),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Error: Could not find product by ID",
			},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			ts.ExecuteTest(t, &testCase)
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	request := helpers.Request{
		Method: http.MethodDelete,
		Url:    "/product/1",
	}

	cases := []helpers.TestCase{
		{
			TestName: "Delete product by id",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Product has been deleted",
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
