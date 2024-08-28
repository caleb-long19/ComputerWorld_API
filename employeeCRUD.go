package main

import (
	"fmt"
	"time"
)

func createNewEmployee() {
	employeeInfo := employeeData{
		employeeName: "",
		employeeRole: "",
	}

	clearData()

	fmt.Println("Please Enter The Employee's Name: ")
	eName := scanUserInput(employeeInfo.employeeName)
	fmt.Println("Product Code: ", eName)

	fmt.Println("Please Enter The Employee's Role: ")
	eRole := scanUserInput(employeeInfo.employeeRole)
	fmt.Println("Product Name: ", eRole)

	storeEmployeeDetails(eName, eRole)

	//store results of employee
	result := databaseCN.Create(&employeeRecords)

	// Check errors and print results to console
	if result.Error != nil {
		panic(result.Error.Error())
	}

	EmployeeInformationApplication()
}

func updateEmployeeRecords() {
	clearData()
	fmt.Println("Please Enter The Name of the Employee you wish to change: ")
	findEmployee = scanUserInput(findEmployee)

	fmt.Println("Please Enter The New Name: ")
	newEmployeeValue = scanUserInput(newEmployeeValue)
	databaseCN.Model(&EmployeeData{}).Select("employee_name").Where("employee_name = ?", findEmployee).Updates(map[string]interface{}{"employee_name": newEmployeeValue})
	fmt.Println("Name has been changed!")

	ProductInformationApplication()
}

func deleteEmployee() {
	selectRecord = ""
	fmt.Println("Deleting Employee Records:")
	selectRecord = scanUserInput(selectRecord)
	fmt.Println("Deleting Employee Data: ", selectRecord)
	databaseCN.Where("Employee_Name = ?", selectRecord).Delete(&EmployeeData{})
	EmployeeInformationApplication()
}

func EmployeeInformationApplication() {
	clearData()
	fmt.Println("Welcome to the Employee Information Page: What would you like to do?")
	fmt.Println("Add Employee, Update Employee, Delete Employee, Exit")
	choosePage = scanUserInput(choosePage)

	fmt.Println("You Chose:", choosePage)

	if choosePage == "Add Employee" {
		fmt.Println("Loading 'Add Employee': Please wait...")
		time.Sleep(1 * time.Second)
		createNewEmployee()
	} else if choosePage == "Update Employee" {
		fmt.Println("Loading 'Update Employee': Please wait...")
		time.Sleep(1 * time.Second)
		updateEmployeeRecords()
	} else if choosePage == "Delete Employee" {
		fmt.Println("Loading 'Delete Employee': Please wait...")
		time.Sleep(1 * time.Second)
		deleteEmployee()
	} else if choosePage == "Exit" {
		fmt.Println("Returning to Home Page: Please wait...")
		time.Sleep(1 * time.Second)
		cwIntroduction()
	}
}
