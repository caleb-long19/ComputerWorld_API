package Console_Application

import (
	"ComputerWorld_API/Model"
)

func storeProductDetails(code, name string, price float64) {
	productRecords = Model.Products{Code: code, Name: name, Price: price}
}

func storeManufacturer(name string) {
	manufacturerRecords = Model.Manufacturer{ManufacturerName: name}
}
