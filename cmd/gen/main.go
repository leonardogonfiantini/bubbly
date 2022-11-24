package main

import (
	"github.com/leonardogonfiantini/bubbly/diagram"
)

func main() {
	

	d := diagram.NewSTR()

	d.CreateFact("SaleFT", []string{"StoreID", "DateID", "ProductID"}, []string{"Quantity", "Receipts"})
	
	
	d.AddDimension("DateDT", []string{"DateID"}, []string{"Date", "Month"}, "SaleFT", "DateID")
	d.AddDimension("StoreDT", []string{"StoreID"}, []string{"Store", "City", "Country", "SalesManager"}, "SaleFT", "StoreID")
	d.AddDimension("ProductDT", []string{"ProductID"}, []string{"Product", "Type", "Category", "Supplier"}, "SaleFT", "ProductID")


	d.RenderDiagram()
}
