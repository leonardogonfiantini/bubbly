package main

import (
	"github.com/leonardogonfiantini/goccia/diagram"
)

func main() {
	

	d := diagram.NewNN()

	InputLayer := d.CreateInputLayer("InputLayer", 2)
	
	HiddenLayer1 := d.CreateHiddenLayer("HiddenLayer1", 4)
	HiddenLayer2 := d.CreateHiddenLayer("HiddenLayer2", 82)
	HiddenLayer3 := d.CreateHiddenLayer("HiddenLayer3", 4)
	
	OutputLayer := d.CreateOutputLayer("OutputLayer", 2)

	d.ConnectLayers(InputLayer, HiddenLayer1)
	d.ConnectLayers(HiddenLayer1, HiddenLayer2)
	d.ConnectLayers(HiddenLayer2, HiddenLayer3)
	d.ConnectLayers(HiddenLayer3, OutputLayer)
	
	d.RenderDiagram()
}
