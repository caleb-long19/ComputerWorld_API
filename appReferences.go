package main

import (
	"bufio"
	"os"
)

// Repeat Application
var choosePage string
var selectRecord string

var findEmployee string
var newEmployeeValue string

var findProduct string
var newProductValue string

// Database
var productRecords StoredProduct
var employeeRecords EmployeeData

func scanUserInput(scanInput string) string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanInput = scanner.Text()

	return scanInput
}
