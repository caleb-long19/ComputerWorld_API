package tests

import (
	"ComputerWorld_API/tests/helpers"
	"fmt"
	"os"
	"testing"
)

var (
	ts *helpers.TestServer
)

func TestMain(m *testing.M) {
	err := os.Setenv("GO_ENV", "test")
	if err != nil {
		fmt.Printf("Error when setting up environment: %v\n", err)
	}

	// Create the test api
	ts = helpers.NewTestServer()
	status := m.Run()
	os.Exit(status)
}
