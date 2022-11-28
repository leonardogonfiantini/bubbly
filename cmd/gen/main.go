package main

import (
	"github.com/leonardogonfiantini/bubbly/diagram"
)

func main() {
	

	d := diagram.NewSTR()

	fact := d.CreateDimension("factDT", []string{"key1", "key2"}, []string{"att1", "att2"}, []string{"red", "blue"}, "green")
	d.RenderDimension(fact)

	dim2 := d.CreateDimension("dim2DT", []string{"key2"}, nil, []string{"red"}, "blue")
	d.RenderDimension(dim2)

	d.JoinDimension(fact, dim2, "key2")


	d.RenderDiagram()
}
