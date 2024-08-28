package Console_Application

import (
	Model2 "ComputerWorld_API/Model"
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
var productRecords Model2.StoredProduct
var employeeRecords Model2.EmployeeData

func scanUserInput(scanInput string) string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanInput = scanner.Text()

	return scanInput
}
