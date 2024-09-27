package tests

import (
	"ComputerWorld_API/db/models"
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

	pd := &models.Product{
		ProductCode:    "DUPLIC8TE",
		ProductName:    "Xbox Duplicate",
		ManufacturerID: 1,
		Stock:          250,
		Price:          400,
	}
	ts.S.Database.Create(pd)

	cases := []helpers.TestCase{
		{
			TestName: "Can create a Product",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Product{
				ProductCode:    "FSDFS3",
				ProductName:    "Xbox Series P",
				ManufacturerID: 1,
				Stock:          250,
				Price:          400,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts:  []string{`"product_name":"Xbox Series P"`},
			},
		},
		{
			TestName: "Can create product when stock is empty!",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Product{
				ProductCode:    "23SDFSF",
				ProductName:    "Xbox Super Cool",
				ManufacturerID: 1,
				Stock:          0,
				Price:          400,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts:  []string{`"product_name":"Xbox Super Cool"`},
			},
		},
		{
			TestName: "Cannot create product as product code contains incorrect formatting",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Product{
				ProductCode:    "DUPLIC8TE@#@#",
				ProductName:    "Xbox Series PS",
				ManufacturerID: 1,
				Stock:          250,
				Price:          400,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot create product as product name contains incorrect formatting",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Product{
				ProductCode:    "DUPLIC8TES",
				ProductName:    "Xbox Series PS##@$",
				ManufacturerID: 1,
				Stock:          250,
				Price:          400,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot create product as product code already exists!",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Product{
				ProductCode:    "DUPLIC8TE",
				ProductName:    "Xbox Series PS",
				ManufacturerID: 1,
				Stock:          250,
				Price:          400,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusConflict,
			},
		},
		{
			TestName: "Cannot create product as product code length is above maximum",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Product{
				ProductCode:    "GGGGGGGGGGGGGGGGGGGGGG",
				ProductName:    "Xbox One Y",
				ManufacturerID: 1,
				Stock:          250,
				Price:          400,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot create product as product name already exists!",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Product{
				ProductCode:    "FSDFSGG3",
				ProductName:    "Xbox Duplicate",
				ManufacturerID: 1,
				Stock:          250,
				Price:          400,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusConflict,
			},
		},
		{
			TestName: "Cannot create product as product name length is above maximum",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Product{
				ProductCode:    "UIGHGJD",
				ProductName:    "XBOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOX",
				ManufacturerID: 1,
				Stock:          250,
				Price:          400,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot create product as product code is empty!",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Product{
				ProductCode:    "",
				ProductName:    "Xbox Series V",
				ManufacturerID: 1,
				Stock:          250,
				Price:          400,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot create product as product name is empty!",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Product{
				ProductCode:    "DSFV3S",
				ProductName:    "",
				ManufacturerID: 1,
				Stock:          250,
				Price:          400,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot create product as manufacturer id is empty!",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Product{
				ProductCode:    "Sony",
				ProductName:    "Xbox Series Z",
				ManufacturerID: 0,
				Stock:          250,
				Price:          400,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot create product as manufacturer id does not exist!",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Product{
				ProductCode:    "TESTXSK",
				ProductName:    "Xbox Series KL",
				ManufacturerID: 99999999,
				Stock:          250,
				Price:          400,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			TestName: "Cannot create product as Price is empty!",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Product{
				ProductCode:    "Sony",
				ProductName:    "Xbox Series Z",
				ManufacturerID: 1,
				Stock:          250,
				Price:          0,
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

func TestGetProduct(t *testing.T) {
	ts.ClearTable("products")

	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/product",
	}

	pd := &models.Product{
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
			TestName: "Can retrieve product by ID as a string was given",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, "1"),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
			},
		},
		{
			TestName: "Can retrieve all products without ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/", request.Url),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
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

	product := &models.Product{
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
			RequestBody: models.Product{
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

	product := &models.Product{
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
