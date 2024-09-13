package tests

import (
	"ComputerWorld_API/db/model"
	"ComputerWorld_API/tests/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestPostProduct(t *testing.T) {
	ts.ClearTable("products")

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/product/",
	}

	cases := []helpers.TestCase{
		{
			TestName: "Can create a Product",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: model.Product{
				ProductCode:    "Sony",
				ProductName:    "Xbox Series Z",
				ManufacturerID: 1,
				Stock:          250,
				Price:          400,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts:  []string{`"product_name":"Xbox Series Z"`},
			},
		},
		{
			TestName: "Cannot create product as it already exists!",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: model.Product{
				ProductCode:    "Sony",
				ProductName:    "Xbox Series Z",
				ManufacturerID: 1,
				Stock:          250,
				Price:          400,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusConflict,
			},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			ts.ExecuteTest(t, &testCase)
		})
	}

}

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
		{
			TestName: "Can retrieve product by id",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, pd.ProductID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					pd.ProductName,
					fmt.Sprintf(`"product_id":%v`, pd.ProductID),
				},
			},
		},
		{
			TestName: "Cannot retrieve product by id",
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

func TestPutProduct(t *testing.T) {
	ts.ClearTable("products")

	request := helpers.Request{
		Method: http.MethodPut,
		Url:    "/product",
	}

	product := &model.Product{
		ProductCode:    "TESTREF",
		ProductName:    "Super Box 360",
		ManufacturerID: 2,
		Stock:          350,
		Price:          400,
	}
	ts.S.Database.Create(product)

	cases := []helpers.TestCase{
		{
			TestName: "Can update Product by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, product.ProductID),
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
				BodyParts:  []string{fmt.Sprintf(`"product_code":"CHZXMG45J"`)},
			},
		},
		{
			TestName: "Cannot update product by ID",
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

func TestDeleteProduct(t *testing.T) {
	ts.ClearTable("products")

	request := helpers.Request{
		Method: http.MethodDelete,
		Url:    "/product",
	}

	product := &model.Product{
		ProductCode:    "TESTREF",
		ProductName:    "TEST_PRODUCT",
		ManufacturerID: 2,
		Stock:          350,
		Price:          400,
	}
	ts.S.Database.Create(product)

	cases := []helpers.TestCase{
		{
			TestName: "Can delete product by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, product.ProductID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
			},
		},
		{
			TestName: "Cannot find and delete product by ID",
			Request: helpers.Request{
				Method: http.MethodDelete,
				Url:    "/product/1000000",
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
