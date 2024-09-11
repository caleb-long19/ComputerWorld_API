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

	Create test data for all tests

	Run through coverage report and make sure you've covered as much as you can
	for each endpoint
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
		// Cannot get manufacturer with invalid ID
		// Can get manufacturer
		{
			TestName: "Retrieve manufacturer by name",
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
	}
	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			ts.ExecuteTest(t, &testCase)
		})
	}
}

func TestDeleteManufacturer(t *testing.T) {

	request := helpers.Request{
		Method: http.MethodDelete,
		Url:    "/manufacturer/1",
	}

	cases := []helpers.TestCase{
		{
			TestName: "Delete manufacturer by name",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Manufacturer has been deleted",
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
