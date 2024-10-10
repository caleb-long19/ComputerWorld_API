package tests

import (
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/tests/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestAdminCreate(t *testing.T) {
	ts.ClearTable("users")

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/admin/",
	}

	admin := &models.Admin{
		Email:    "johnadmin@gmail.com",
		Name:     "John Test",
		Password: "TestPass15!",
	}
	ts.S.Database.Create(admin)

	cases := []helpers.TestCase{
		{
			TestName: "Can create am admin",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Admin{
				Email:    "sarahadmin@gmail.com",
				Name:     "Sarah Admin",
				Password: "TestPass12!",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts:  []string{`"name":"Sarah Admin"`},
			},
		},
		{
			TestName: "Cannot create admin with a duplicate email",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Admin{
				Email:    "johnadmin@gmail.com",
				Name:     "John Admin",
				Password: "TestPass12!",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot create admin with an incorrect email format",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Admin{
				Email:    "kelly_email_admin@gmail",
				Name:     "Kelly Admin",
				Password: "TestPass12!",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot create admin with an invalid password",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.Admin{
				Email:    "jackadmin@gmail.com",
				Name:     "Jack Admin",
				Password: "password",
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

func TestAdminGet(t *testing.T) {
	ts.ClearTable("admins")

	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/admin",
	}

	admin := &models.Admin{
		Email:    "johnadmin@gmail.com",
		Name:     "John Admin",
		Password: "TestPass15!",
	}
	ts.S.Database.Create(admin)

	cases := []helpers.TestCase{
		{
			TestName: "Can get admin by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, admin.AdminID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					admin.Name,
					fmt.Sprintf(`"admin_id":%v`, admin.AdminID),
				},
			},
		},
		{
			TestName: "Can retrieve all admins without ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/", request.Url),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
			},
		},
		{
			TestName: "Cannot get admin that does not exist",
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

func TestAdminUpdate(t *testing.T) {
	ts.ClearTable("admins")

	request := helpers.Request{
		Method: http.MethodPut,
		Url:    "/admin",
	}

	admin := &models.Admin{
		Email:    "johnadmin@gmail.com",
		Name:     "John Admin",
		Password: "TestPass15!",
	}
	ts.S.Database.Create(admin)

	cases := []helpers.TestCase{
		{
			TestName: "Can update admin by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, admin.AdminID),
			},
			RequestBody: models.User{
				Email:    "johnadminnew@gmail.com",
				Name:     "Johnny Admin",
				Password: "TestPass17!!",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts:  []string{fmt.Sprintf(`"email":"johnadminnew@gmail.com"`)},
			},
		},
		{
			TestName: "Cannot update admin as email format was incorrect",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, admin.AdminID),
			},
			RequestBody: models.User{
				Email:    "admin_test.gmail.com",
				Name:     "John Testing",
				Password: "TestPass17!!",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot update admin as name has special characters",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, admin.AdminID),
			},
			RequestBody: models.User{
				Email:    "jackadminnew@gmail.com",
				Name:     "Jack Admin@@@",
				Password: "TestPass17!!",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot update admin as password has invalid format",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, admin.AdminID),
			},
			RequestBody: models.User{
				Email:    "sarahadminnew@gmail.com",
				Name:     "Sarah Admin",
				Password: "Testing",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot update admin as they do not exist",
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

func TestAdminDelete(t *testing.T) {
	request := helpers.Request{
		Method: http.MethodDelete,
		Url:    "/admin",
	}

	admin := &models.Admin{
		Email:    "deletetest@gmail.com",
		Name:     "Delete Test",
		Password: "TestPass15!",
	}
	ts.S.Database.Create(admin)

	cases := []helpers.TestCase{
		{
			TestName: "Can delete admin by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, admin.AdminID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
			},
		},
		{
			TestName: "Cannot find and delete admin that does not exist",
			Request: helpers.Request{
				Method: http.MethodDelete,
				Url:    "/admin/10000000",
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
