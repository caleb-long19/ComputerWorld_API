package helpers

import (
	"io"
)

type TestCase struct {
	TestName           string
	Request            Request
	RequestContentType string
	RequestBody        io.Reader
	ExpectedStatusCode int
	ExpectedBody       string
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
