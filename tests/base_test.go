package tests

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := os.Setenv("GO_ENV", "test")

	if err != nil {
		fmt.Printf("Error when setting up environment: %v\n", err)
	}

	status := m.Run()

	os.Exit(status)
}
