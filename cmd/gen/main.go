package main

import (
	"github.com/leonardogonfiantini/bubbly/diagram"
)

func main() {
	

	d := diagram.NewDFM()

	d.CreateFact("SALE", "quantity receipts numberOfCustomers unitPrice")

	d.AddSequenceDimension("store storeCity state country", "SALE")
	d.AddSequenceDimension("date month quarter", "SALE")
	d.AddSequenceDimension("product type", "SALE")

	d.AddDimension("salesManager", "store")
	d.AddDimension("salesDistrict", "store")
	d.AddSequenceDescriptive("address telephone", "store")

	d.AddDimension("day", "date")
	d.AddDimension("week", "date")
	d.AddDimension("holiday", "date")
	d.AddDimension("year", "quarter")

	d.AddOptional("diet", "product")
	d.AddDescriptive("weight", "product")
	d.AddSequenceDimension("brand brandCity", "product")
	
	d.AddSequenceDimension("category department", "type")
	d.AddDescriptive("departmentHead", "department")
	d.AddDimension("marketingGroup", "type")
	d.AddDescriptive("director", "marketingGroup")

	d.RenderDiagram()
}
