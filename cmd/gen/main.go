package main

import (
	"log"
	"os"

	"github.com/leonardogonfiantini/bubbly/diagram"
)

func main() {
	

	d := diagram.NewSTR()
	graph := d.Graph


	d.CreateFact("SaleFT", []string{"StoreID", "DateID", "ProductID"}, []string{"Quantity", "Receipts"})
	
	
	d.AddDimension("DateDT", []string{"DateID"}, []string{"Date", "Month"}, "SaleFT", "DateID")
	d.AddDimension("StoreDT", []string{"StoreID"}, []string{"Store", "City", "Country", "SalesManager"}, "SaleFT", "StoreID")
	d.AddDimension("ProductDT", []string{"ProductID"}, []string{"Product", "Type", "Category", "Supplier"}, "SaleFT", "ProductID")


	output := graph.String()


	f, err := os.Create("dot.dot")
	if err != nil {
        log.Fatal(err)
    }

	defer f.Close()
	f.WriteString(output)

}