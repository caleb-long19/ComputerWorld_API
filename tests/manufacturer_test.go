package tests

import (
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/tests/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestPostManufacturer(t *testing.T) {
	ts.ClearTable("manufacturers")

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/manufacturer/",
	}

	mf := &models.Manufacturer{
		ManufacturerName: "Microsoft",
	}
	ts.S.Database.Create(mf)

	cases := []helpers.TestCase{
		{
			TestName: "Can create a manufacturer",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Manufacturer{
				ManufacturerName: "Wacom",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts:  []string{`"manufacturer_name":"Wacom"`},
			},
		},
		{
			TestName: "Cannot create manufacturer as they already exists!",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Manufacturer{
				ManufacturerName: "Microsoft",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusConflict,
			},
		},
		{
			TestName: "Cannot create manufacturer as the name exceeds the maximum length!",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Manufacturer{
				ManufacturerName: "Super Amazing Microsoft Dream Team",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusLengthRequired,
			},
		},
		{
			TestName: "Cannot create manufacturer as no name was given",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Manufacturer{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotAcceptable,
			},
		},
		{
			TestName: "Cannot create manufacturer as name was left blank",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Manufacturer{
				ManufacturerName: "",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotAcceptable,
			},
		},
		{
			TestName: "Cannot create manufacturer as name contains special characters",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Manufacturer{
				ManufacturerName: "Test####@",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotAcceptable,
			},
		},
		{
			TestName: "Cannot create manufacturer when an ID is given",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Manufacturer{
				ManufacturerID: 1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotAcceptable,
			},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.TestName, func(t *testing.T) {
			ts.ExecuteTest(t, &testCase)
		})
	}

}

func TestGetManufacturer(t *testing.T) {
	ts.ClearTable("manufacturers")

	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/manufacturer",
	}

	mf := &models.Manufacturer{
		ManufacturerName: "Microsoft",
	}
	ts.S.Database.Create(mf)

	cases := []helpers.TestCase{
		{
			TestName: "Can retrieve manufacturer by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, mf.ManufacturerID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					mf.ManufacturerName,
					fmt.Sprintf(`"manufacturer_id":%v`, mf.ManufacturerID),
				},
			},
		},
		{
			TestName: "Can retrieve manufacturer by ID as a string was given",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, "1"),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
			},
		},
		{
			TestName: "Cannot retrieve manufacturer by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, 10000),
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

func TestPutManufacturer(t *testing.T) {
	ts.ClearTable("manufacturers")

	request := helpers.Request{
		Method: http.MethodPut,
		Url:    "/manufacturer",
	}

	mf := &models.Manufacturer{
		ManufacturerName: "Microsoft",
	}
	ts.S.Database.Create(mf)

	cases := []helpers.TestCase{
		{
			TestName: "Can update Manufacturer by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, mf.ManufacturerID),
			},
			RequestBody: models.Manufacturer{
				ManufacturerName: "Akira",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts:  []string{fmt.Sprintf(`"manufacturer_name":"Akira"`)},
			},
		},
		{
			TestName: "Can update Manufacturer by ID and include numbers in the name",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, mf.ManufacturerID),
			},
			RequestBody: models.Manufacturer{
				ManufacturerName: "AkiraTest123",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts:  []string{fmt.Sprintf(`"manufacturer_name":"AkiraTest123"`)},
			},
		},
		{
			TestName: "Cannot update manufacturer as no body was given",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, mf.ManufacturerID),
			},
			RequestBody: models.Manufacturer{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotAcceptable,
			},
		},
		{
			TestName: "Cannot update manufacturer as no name was given",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, mf.ManufacturerID),
			},
			RequestBody: models.Manufacturer{
				ManufacturerName: "",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotAcceptable,
			},
		},
		{
			TestName: "Cannot update manufacturer as name was special characters",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, mf.ManufacturerID),
			},
			RequestBody: models.Manufacturer{
				ManufacturerName: "MicrosoftTest#####",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotAcceptable,
			},
		},
		{
			TestName: "Cannot update manufacturer as name exceeded maximum length",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, mf.ManufacturerID),
			},
			RequestBody: models.Manufacturer{
				ManufacturerName: "MicrosoftTestingToMakeSureTheLengthIsNotTooMuch",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusLengthRequired,
			},
		},
		{
			TestName: "Cannot update manufacturer by ID",
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

func TestDeleteManufacturer(t *testing.T) {
	ts.ClearTable("manufacturers")

	request := helpers.Request{
		Method: http.MethodDelete,
		Url:    "/manufacturer",
	}

	mf := &models.Manufacturer{
		ManufacturerName: "Microsoft",
	}
	ts.S.Database.Create(mf)

	cases := []helpers.TestCase{
		{
			TestName: "Can delete manufacturer by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, mf.ManufacturerID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
			},
		},
		{
			TestName: "Cannot find and delete manufacturer by ID",
			Request: helpers.Request{
				Method: http.MethodDelete,
				Url:    "/manufacturer/10000000",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			TestName: "Cannot delete manufacturer as no ID was given",
			Request: helpers.Request{
				Method: request.Method,
				Url:    "/manufacturer/",
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
