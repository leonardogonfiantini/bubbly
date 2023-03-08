package diagram

import (
	"github.com/awalterschulze/gographviz"

	"os"
	"log"
	"strconv"
)

/*
Struct NN defines the neural network schema and includes a pointer to 
a graph.
*/
type NN struct {
	Graph *gographviz.Graph
}

/*
Attributes are maps of attributes for modifying nodes and edges 
in the nn schema.
*/
var (
	NN_nodeAtt = map[string]string{"shape":"circle", "label":"\"\""}
)

/*
Function NewNN creates a new NN object for creating the nn schema.
*/
func NewNN() *NN {

	graphAst, _ := gographviz.ParseString(`digraph G {
		layout="circo"
	}`)

	graph := gographviz.NewGraph()
	if err := gographviz.Analyse(graphAst, graph); err != nil {
		panic(err)
	}

	return &NN {
		graph,
	}
}

/*
Function CreateInputLayer, adds the input layer in the nn schema
*/
func (schema *NN) CreateInputLayer(num_inputs int) {

	for i := 0; i < num_inputs; i++ {
		schema.Graph.AddNode("G", strconv.Itoa(i), NN_nodeAtt)
	}

}

/*
Function RenderDiagram renders the entire schema to a dot file
*/
func (schema *NN) RenderDiagram() {
	output := schema.Graph.String()

	f, err := os.Create("dot.dot")
	if err != nil {
        log.Fatal(err)
    }

	defer f.Close()
	f.WriteString(output)
}


