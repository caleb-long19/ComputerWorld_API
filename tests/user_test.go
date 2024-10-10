package tests

import (
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/tests/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestUserCreate(t *testing.T) {
	ts.ClearTable("users")

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/user/",
	}

	user := &models.User{
		Email:    "johntest@gmail.com",
		Name:     "John Test",
		Password: "TestPass15!",
	}
	ts.S.Database.Create(user)

	cases := []helpers.TestCase{
		{
			TestName: "Can create a user",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.User{
				Email:    "sarahtest@gmail.com",
				Name:     "Sarah Test",
				Password: "TestPass12!",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts:  []string{`"name":"Sarah Test"`},
			},
		},
		{
			TestName: "Cannot create user with a duplicate email",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.User{
				Email:    "johntest@gmail.com",
				Name:     "John Test",
				Password: "TestPass12!",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot create user with an incorrect email format",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.User{
				Email:    "kelly_email@gmail",
				Name:     "Kelly Test",
				Password: "TestPass12!",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot create user with an invalid password",
			Request: helpers.Request{
				Method: request.Method,
				Url:    request.Url,
			},
			RequestBody: models.User{
				Email:    "jacktest@gmail.com",
				Name:     "Jack Test",
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

func TestUserGet(t *testing.T) {
	ts.ClearTable("users")

	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/user",
	}

	user := &models.User{
		Email:    "johntest@gmail.com",
		Name:     "John Test",
		Password: "TestPass15!",
	}
	ts.S.Database.Create(user)

	cases := []helpers.TestCase{
		{
			TestName: "Can get user by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, user.UserID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					user.Name,
					fmt.Sprintf(`"user_id":%v`, user.UserID),
				},
			},
		},
		{
			TestName: "Can retrieve all users without ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/", request.Url),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
			},
		},
		{
			TestName: "Cannot get user that does not exist",
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

func TestUserUpdate(t *testing.T) {
	ts.ClearTable("users")

	request := helpers.Request{
		Method: http.MethodPut,
		Url:    "/user",
	}

	user := &models.User{
		Email:    "johntest@gmail.com",
		Name:     "John Test",
		Password: "TestPass15!",
	}
	ts.S.Database.Create(user)

	cases := []helpers.TestCase{
		{
			TestName: "Can update user by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, user.UserID),
			},
			RequestBody: models.User{
				Email:    "johntestnew@gmail.com",
				Name:     "Johnny Testing",
				Password: "TestPass17!!",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts:  []string{fmt.Sprintf(`"email":"johntestnew@gmail.com"`)},
			},
		},
		{
			TestName: "Cannot update user as email format was incorrect",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, user.UserID),
			},
			RequestBody: models.User{
				Email:    "test.gmail.com",
				Name:     "John Testing",
				Password: "TestPass17!!",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot update user as name has special characters",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, user.UserID),
			},
			RequestBody: models.User{
				Email:    "jacktestnew@gmail.com",
				Name:     "Jack Testing@@@",
				Password: "TestPass17!!",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot update user as password has invalid format",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, user.UserID),
			},
			RequestBody: models.User{
				Email:    "sarahtestnew@gmail.com",
				Name:     "Sarah Test",
				Password: "Testing",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "Cannot update user as they do not exist",
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

func TestUserDelete(t *testing.T) {
	request := helpers.Request{
		Method: http.MethodDelete,
		Url:    "/user",
	}

	user := &models.User{
		Email:    "deletetest@gmail.com",
		Name:     "Delete Test",
		Password: "TestPass15!",
	}
	ts.S.Database.Create(user)

	cases := []helpers.TestCase{
		{
			TestName: "Can delete user by ID",
			Request: helpers.Request{
				Method: request.Method,
				Url:    fmt.Sprintf("%v/%v", request.Url, user.UserID),
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
			},
		},
		{
			TestName: "Cannot find and delete user that does not exist",
			Request: helpers.Request{
				Method: http.MethodDelete,
				Url:    "/user/10000000",
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
