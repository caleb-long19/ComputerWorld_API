package main

func storeProductDetails(code, name string, price float64) {
	productRecords = StoredProduct{Code: code, Name: name, Price: price}
}

func storeEmployeeDetails(name, role string) {
	employeeRecords = EmployeeData{EmployeeName: name, EmployeeRole: role}
}
