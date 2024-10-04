package tests

import (
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/tests/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestManufacturerCreate(t *testing.T) {
	ts.ClearTable("manufacturers")

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/manufacturer/",
	}

	mf := &models.Manufacturer{
		ManufacturerName: "Microsoftest",
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
				ManufacturerName: "Microsoftest",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
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
				StatusCode: http.StatusBadRequest,
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
				StatusCode: http.StatusBadRequest,
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

func TestManufacturerGet(t *testing.T) {
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
			TestName: "Can get manufacturer by ID",
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
			TestName: "Can retrieve all manufacturers without ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/", request.Url),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
			},
		},
		{
			TestName: "Cannot get manufacturer that does not exist",
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

func TestManufacturerUpdate(t *testing.T) {
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
			TestName: "Cannot update manufacturer as no name was given",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, mf.ManufacturerID),
			},
			RequestBody: models.Manufacturer{
				ManufacturerName: "",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot update manufacturer as name has special characters",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, mf.ManufacturerID),
			},
			RequestBody: models.Manufacturer{
				ManufacturerName: "MicrosoftTest#####",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot update manufacturer as they do not exist",
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

func TestManufacturerDelete(t *testing.T) {
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
			TestName: "Cannot find and delete manufacturer that does not exist",
			Request: helpers.Request{
				Method: http.MethodDelete,
				Url:    "/manufacturer/10000000",
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
