package main

import (
	"github.com/leonardogonfiantini/bubbly/diagram"
)

func main() {
	

	d := diagram.NewSTR()

	d.CreateFact("SaleFT", []string{"StoreID", "DateID", "ProductID"}, []string{"Quantity", "Receipts"}, []string{"yellow", "green", "red"}, "lime")
	
	d.AddDimension("DateDT", []string{"DateID"}, []string{"Date", "Month"}, "SaleFT", "DateID", []string{"green"}, "green")
	d.AddDimension("StoreDT", []string{"StoreID"}, []string{"Store", "City", "Country", "SalesManager"}, "SaleFT", "StoreID", []string{"yellow"}, "yellow")
	d.AddDimension("ProductDT", []string{"ProductID"}, []string{"Product", "Type", "Category", "Supplier"}, "SaleFT", "ProductID", []string{"red"}, "red")

	d.RenderDiagram()
}
