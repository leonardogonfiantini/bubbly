package main

import (
	"github.com/leonardogonfiantini/goccia/diagram"
)

func main() {
	

	d := diagram.NewNN()

	d.CreateInputLayer(4)

	d.RenderDiagram()
}
