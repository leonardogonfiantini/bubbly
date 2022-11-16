package main

import (
	"log"
	"os"

	"github.com/leonardogonfiantini/bubbly/diagram"
)

func main() {
	

	d := diagram.NewDFM()
	graph := d.Graph

	d.CreateFact("SALE", []string{"quantity", "receipts", "unitPrice", "numberOfCustomers"})

	d.AddNode("product", "SALE")
	d.AddNode("date", "SALE")
	d.AddNode("store", "SALE")

	d.AddNode("brand", "product")
	d.AddNode("brandCity", "brand")
	
	d.AddNode("type", "product")
	d.AddNode("marketingGroup", "type")
	d.AddNode("category", "type")
	d.AddNode("department", "category")

	d.AddSequenceDescriptive([]string{"desc1", "desc2", "desc3", "desc4"}, "brand")

	d.AddHierarchy([]string{"prova1", "prova2"}, "SALE", "prova")


	d.AddSequenceNode([]string{"month","quarter","year"}, "date")
	d.AddNode("holiday", "date")
	d.AddNode("day", "date")
	d.AddNode("week", "date")


	d.AddSequenceNode([]string{"storeCity", "state", "country"}, "store")
	d.AddOptional("salesManager", "store")
	d.AddNode("salesDistrict", "store")

	output := graph.String()


	f, err := os.Create("dot.dot")
	if err != nil {
        log.Fatal(err)
    }

	defer f.Close()
	f.WriteString(output)

}