package diagram

import (
	"github.com/awalterschulze/gographviz"

	"log"
	"os"
	"strconv"
)

/*
Struct NN defines the neural network schema and includes a pointer to
a graph.
*/
type NN struct {
	Graph *gographviz.Graph
	hiddens_layer int
}

type Layer struct {
	name string
	nums_node int
}

/*
Attributes are maps of attributes for modifying nodes and edges 
in the nn schema.
*/
var (
	NN_nodeAtt = map[string]string{"shape":"circle", "label":"\"\""}
	NN_connectAtt = map[string]string{"arrowhead":"none"}
)

/*
Function NewNN creates a new NN object for creating the nn schema.
*/
func NewNN() *NN {

	graphAst, _ := gographviz.ParseString(`digraph G {
	}`)

	graph := gographviz.NewGraph()
	if err := gographviz.Analyse(graphAst, graph); err != nil {
		panic(err)
	}

	return &NN {
		graph,
		0,
	}
}

/*
Function CreateInputLayer, adds the input layer in the nn schema
*/
func (schema *NN) CreateInputLayer(name string, nums_node int) *Layer {

	if name == "" {
		name = "InputLayer"
	}

	if err := schema.Graph.AddSubGraph("G", name, nil); err != nil {
		panic(err)
	}

	for i := 0; i < nums_node; i++ {	
		if err := schema.Graph.AddNode(name, name+strconv.Itoa(i), NN_nodeAtt); err != nil {
			panic(err)
		}

		if err := schema.Graph.AddNode(name, "HInput"+strconv.Itoa(i), map[string]string{"style":"invis"}); err != nil {
			panic(err)
		}

		if err := schema.Graph.AddEdge("HInput"+strconv.Itoa(i), name+strconv.Itoa(i), true, nil); err != nil {
			panic(err)
		}
	}

	return &Layer{
		name,
		nums_node,
	}
}

/*
Function CreateHiddenLayer, adds an hidden layer in the nn schema
*/
func (schema *NN) CreateHiddenLayer(name string, nums_node int) *Layer {

	schema.hiddens_layer++

	if name == "" {
		name = "HiddenLayer" + strconv.Itoa(schema.hiddens_layer)
	}

	if err := schema.Graph.AddSubGraph("G", name, nil); err != nil {
		panic(err)
	}

	for i := 0; i < nums_node; i++ {
		if err := schema.Graph.AddNode(name, name+strconv.Itoa(i), NN_nodeAtt); err != nil {
			panic(err)
		}
	}

	return &Layer {
		name,
		nums_node,
	}
}

/*
Function CreateOutputLayer, creates the output layer in the nn schema
*/
func (schema *NN) CreateOutputLayer(name string, nums_node int) *Layer {
	if name == "" {
		name = "OutputLayer"
	}

	if err := schema.Graph.AddSubGraph("G", name, nil); err != nil {
		panic(err)
	}

	for i := 0; i < nums_node; i++ {	
		if err := schema.Graph.AddNode(name, name+strconv.Itoa(i), NN_nodeAtt); err != nil {
			panic(err)
		}

		if err := schema.Graph.AddNode(name, "HOutput"+strconv.Itoa(i), map[string]string{"style":"invis"}); err != nil {
			panic(err)
		}

		if err := schema.Graph.AddEdge(name+strconv.Itoa(i), "HOutput"+strconv.Itoa(i), true, nil); err != nil {
			panic(err)
		}
	}

	return &Layer{
		name,
		nums_node,
	}
}

/*
Function ConnectLayers, connect two layers
*/
func (schema *NN) ConnectLayers(src *Layer, dst *Layer) {
	for i := 0; i < src.nums_node; i++ {
		for j := 0; j < dst.nums_node; j ++ {
			schema.Graph.AddEdge(src.name+strconv.Itoa(i), dst.name+strconv.Itoa(j), true, NN_connectAtt)
		}
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


