package main

import (
	"log"
	"os"

	"github.com/leonardogonfiantini/bubbly/diagram"
)

func main() {
	

	d := diagram.NewDFM()
	graph := d.Graph

	d.CreateFact("provaciao", []string{"ciao1", "ciao2", "ciao3"})

	d.AddNode("ciao", "fact")
	d.AddNode("bella", "fact")
	d.AddNode("ok", "ciao")

	d.AddNode("3", "fact")
	d.AddNode("4", "fact")

	d.AddNode("1", "ok")
	d.AddNode("2", "ok")

	d.AddNode("88", "g")
	d.AddNode("41", "ciao")

	d.AddOptional("32", "ciao")

	d.AddConvergence("g", "2")
	d.AddConvergence("g", "1")

	d.AddDescriptive("ciao", "scialla")



	output := graph.String()


	f, err := os.Create("dot.dot")
	if err != nil {
        log.Fatal(err)
    }

	defer f.Close()
	f.WriteString(output)

}