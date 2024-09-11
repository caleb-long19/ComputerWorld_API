package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCase struct {
	TestName           string
	Request            Request
	RequestContentType string
	RequestReader      io.Reader
	RequestBody        interface{}
	ExpectedStatusCode int
	ExpectedBody       string
	Expected           ExpectedResponse
	Setup              func(testCase *TestCase)
	DisplayResponse    bool
}

type PathParam struct {
	Name  string
	Value string
}

type Request struct {
	Method string
	Url    string
	Path   *PathParam
}

type ExpectedResponse struct {
	StatusCode int
	BodyPart   string
	BodyParts  []string
}

func (ts *TestServer) ExecuteTest(t *testing.T, testCase *TestCase) {
	// Run any setup required before we execute the request
	if testCase.Setup != nil {
		testCase.Setup(testCase)
	}
	req, err := ts.GenerateRequest(testCase)
	if err != nil {
		t.Fatalf("unable to Generate Request: %v", err)
	}
	res := ts.ExecuteRequest(req)
	ts.ValidateResults(t, testCase, res)
}

func (ts *TestServer) ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	// Create a new recorder then process request with server.
	rr := httptest.NewRecorder()
	ts.S.Echo.ServeHTTP(rr, req)
	return rr
}

func (ts *TestServer) GenerateRequest(testCase *TestCase) (*http.Request, error) {
	reqJson, err := json.Marshal(testCase.RequestBody)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	if testCase.RequestReader != nil {
		req, err = http.NewRequest(testCase.Request.Method, testCase.Request.Url, testCase.RequestReader)
	} else {
		req, err = http.NewRequest(testCase.Request.Method, testCase.Request.Url, bytes.NewBuffer(reqJson))
	}

	if err != nil {
		return nil, err
	}

	if testCase.RequestContentType != "" {
		req.Header.Set(echo.HeaderContentType, testCase.RequestContentType)
	} else {
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	}

	return req, nil
}

func (ts *TestServer) ValidateResults(t *testing.T, test *TestCase, res *httptest.ResponseRecorder) {
	if test.DisplayResponse {
		fmt.Println("Request Output: ")
		fmt.Println(res.Body.String())
	}

	if res.Code != 0 {
		assert.Equal(t, test.Expected.StatusCode, res.Code)
	}

	if test.Expected.BodyPart != "" {
		assert.Contains(t, res.Body.String(), test.Expected.BodyPart)
	}
	if len(test.Expected.BodyParts) > 0 {
		for _, expectedText := range test.Expected.BodyParts {
			assert.Contains(t, res.Body.String(), expectedText)
		}
	}
}

func (ts *TestServer) ClearTable(tableName string) {
	err := ts.S.Database.Exec(fmt.Sprintf("DELETE FROM %v", tableName)).Error
	if err != nil {
		log.Fatalf("You can't clear that table. Err: %v", err)
	}

	// TODO: CAUSING SOME PROBLEMS WITH SQLITE - NOT NEEDED AT THE MOMENT
	//err = ts.S.Database.Exec(fmt.Sprintf("ALTER TABLE %v AUTO_INCREMENT = 1", tableName)).Error
	//if err != nil {
	//	log.Fatalf("Error setting autoincrement. Err: %v", err)
	//}
}
