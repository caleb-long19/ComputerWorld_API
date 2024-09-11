package tests

import (
	"ComputerWorld_API/database/model"
	"ComputerWorld_API/tests/helpers"
	"fmt"
	"net/http"
	"testing"
)

// TODO: Add in some helpers to clear database tables
/*
	Change controllers to be like manufacturer one - Done

	Add in ability to clear tables for tests - Done

	Create test data for all tests - Done

	Run through coverage report and make sure you've covered as much as you can for each endpoint
*/

func TestGetManufacturer(t *testing.T) {
	ts.ClearTable("manufacturers")

	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/manufacturer",
	}

	mf := &model.Manufacturer{
		ManufacturerName: "Microsoft",
	}
	ts.S.Database.Create(mf)

	cases := []helpers.TestCase{
		// Cannot get manufacturer with invalid ID - Done
		// Can get manufacturer - Done
		{
			TestName: "Test 1 - Retrieve manufacturer by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, mf.ManufacturerID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf("manufacturer_id %v", mf.ManufacturerID),
				},
			},
		},
		{
			TestName: "Test 2 - Error 404: Fail to retrieve manufacturer by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, 10000),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Could not find manufacturer by that ID",
			},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			ts.ExecuteTest(t, &testCase)
		})
	}
}

//func TestPostManufacturer(t *testing.T) {
//	ts.ClearTable("manufacturers")
//
//	request := helpers.Request{
//		Method: http.MethodPost,
//		Url:    "/manufacturer",
//	}
//
//	cases := []helpers.TestCase{
//		{
//			TestName: "Test 1 - Create Manufacturer",
//			Request: helpers.Request{
//				Method: request.Method,
//				Url:    request.Url,
//			},
//			Expected: helpers.ExpectedResponse{
//				StatusCode: http.StatusCreated,
//				BodyPart:   "Manufacturer created successfully",
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

func TestPutManufacturer(t *testing.T) {
	ts.ClearTable("manufacturers")

	request := helpers.Request{
		Method: http.MethodPut,
		Url:    "/manufacturer/2",
	}

	cases := []helpers.TestCase{
		{
			TestName: "Test 1 - Update Manufacturer by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: model.Manufacturer{
				ManufacturerName: "Akira",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Manufacturer updated successfully",
			},
		},
		{
			TestName: "Test 2 - Error 404: Fail to update manufacturer by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, 100000),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Error: Could not find manufacturer by that ID",
			},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			ts.ExecuteTest(t, &testCase)
		})
	}
}

func TestDeleteManufacturer(t *testing.T) {
	ts.ClearTable("manufacturers")

	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/manufacturer/3",
	}

	cases := []helpers.TestCase{
		{
			TestName: "Test 1 - Delete manufacturer by ID",
			Request: helpers.Request{
				Method: http.MethodDelete,
				Url:    request.Url,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Success: Manufacturer has been deleted",
			},
		},
		{
			TestName: "Test 2 - Error 404: Fail to find and delete Manufacturer by ID",
			Request: helpers.Request{
				Method: http.MethodDelete,
				Url:    "/manufacturer/1000",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Error: Could not find manufacturer by that ID",
			},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			ts.ExecuteTest(t, &testCase)
		})
	}
}
