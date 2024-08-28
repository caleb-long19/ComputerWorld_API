package Console_Application

import (
	"ComputerWorld_API/Model"
)

func storeProductDetails(code, name string, price float64) {
	productRecords = Model.StoredProduct{Code: code, Name: name, Price: price}
}

func storeEmployeeDetails(name, role string) {
	employeeRecords = Model.EmployeeData{EmployeeName: name, EmployeeRole: role}
}
