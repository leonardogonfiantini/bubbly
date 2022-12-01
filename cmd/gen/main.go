package main

import (
	"github.com/leonardogonfiantini/bubbly/diagram"
)

func main() {
	

	d := diagram.NewDFM()

	d.CreateFact("factTable", "att1 att2 att3 att4")

	d.AddDimension("dimension1", "factTable")
	d.AddDimension("dimension2", "dimension1")
	d.AddDimension("dimension3", "dimension2")
	d.AddDimension("dimension4", "dimension2")
	d.AddDimension("dimension5", "dimension4")
	d.AddDimension("dimension6", "dimension4")


	d.RenderDiagram()
}
