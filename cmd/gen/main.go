package main

import (
	"github.com/leonardogonfiantini/bubbly/diagram"
)

func main() {
	

	d := diagram.NewSTR()

	fact := d.CreateDimension("factDT", "key1 key2", "att1 att2")
	dim2 := d.CreateDimension("dim2DT", "key2 key4", "keys ok ok ok ")

	d.JoinDimension(fact, dim2, "key2")

	d.RenderDiagram()
}
