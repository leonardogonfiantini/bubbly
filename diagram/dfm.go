package diagram

import (
	"github.com/awalterschulze/gographviz"

	"strings"
	"log"
	"os"
)

/*
Struct for DFM schema
*/
type Dfm struct {
	Graph *gographviz.Graph	 //graph
}

type Fact struct {
	name string
	attributes []string
}

/*
Attributes for modify nodes in the DFM schema
*/
var  (
	DFM_nodeAtt = map[string]string{"shape":"circle"}
	DFM_edgeAtt = map[string]string{"arrowhead":"none", "len":"0.5"}
	DFM_factAtt =  map[string]string{"shape":"plain", "root":"true"}
	DFM_descriptiveAtt = map[string]string{"shape":"underline"}
	DFM_optionalAtt = map[string]string{"arrowhead":"icurve"}
	DFM_hierarchyAtt = map[string]string{"arrowhead":"none"}
)

/*
Create a new object *Dfm to use for creating the dfm schema
*/
func NewDFM() *Dfm {

	graphAst, _ := gographviz.ParseString(`digraph G { 
		layout=twopi;
		ranksep=2.75;
	}`)
	graph := gographviz.NewGraph()
	if err := gographviz.Analyse(graphAst, graph); err != nil {
		panic(err)
	}

	return &Dfm{
		graph,
	}
}

/*
Create an object *Fact to use for create a fact table in the schema
*/
func (schema *Dfm) CreateFact(title string, attributes string) *Fact {

	t_attributes := strings.Split(attributes, " ")

	fact := &Fact {
		title,
		t_attributes,
	}


	schema.RenderFact(fact)
	return fact
}

/*
Render the fact
*/
func (schema *Dfm) RenderFact(f *Fact) {

	label := `<<table border="0" cellborder="1" cellspacing="0" cellpadding="20"> <tr> <td bgcolor="lightblue">`+f.name+`</td> </tr>`
	for _, att := range(f.attributes) {
		label += `<tr> <td>`+att+`</td> </tr>`
	}
	label += `</table>>`

	fact_att := DFM_factAtt
	fact_att["label"] = label

	schema.Graph.AddNode("G", f.name, fact_att)
}

/*
Add a node with label = label to a node with label = attach
*/
func (schema *Dfm) AddDimension(label string, attach string) {
	
	//refactor this
	node_att := DFM_nodeAtt
	node_att["xlabel"] = label
	node_att["label"] = "\"\""
	node_att["fixedsize"] = "true"

	schema.Graph.AddNode("G", label, node_att)
	schema.Graph.AddEdge(attach, label, true, DFM_edgeAtt)
}


/*
Add multiple nodes with label = labels, starting from node with label = startAttach to the node with label = labels[len(labels)]
*/
func (schema *Dfm) AddSequenceDimension(labels string, startAttach string) {

	arr_labels := strings.Split(labels, " ")

	schema.AddDimension(arr_labels[0], startAttach)
	for i := 0; i < len(arr_labels)-1; i++ {
		schema.AddDimension(arr_labels[i+1], arr_labels[i])
	}

}

/*
Add a convergence node with label = label to a node with label = attach
*/
func (schema *Dfm) AddConvergence(label string, attach string) {
	node_att := DFM_nodeAtt
	node_att["label"] = label

	schema.Graph.AddNode("G", label, node_att)
	schema.Graph.AddEdge(attach, label, true, nil)
}

/*
Add a hierachy with 2 or more node with label = labels starting from node with label = to
*/
func (schema *Dfm) AddHierarchy(labels string, from string, to string) {

	node_att := DFM_nodeAtt
	node_att["label"] = to

	schema.Graph.AddNode("G", to, node_att)

	for _, label := range strings.Split(labels, " ") {
		tmpAtt := DFM_hierarchyAtt
		tmpAtt["label"] = label
		schema.Graph.AddEdge(from, to, true, tmpAtt)
	}

}

/*
Add an optional node with label = label starting from a node with label = attach
*/
func (schema *Dfm) AddOptional(label string, attach string) {

	node_att := DFM_nodeAtt
	node_att["label"] = label

	schema.Graph.AddNode("G", label, node_att)
	schema.Graph.AddEdge(attach, label, true, DFM_optionalAtt)
}

/*
Add a new descriptive attribute with the label = label to a node with label = to
*/
func (schema *Dfm) AddDescriptive(label string, to string) {

	schema.Graph.AddNode("G", label, DFM_descriptiveAtt)
	schema.Graph.AddEdge(to, label, true, DFM_edgeAtt)
}

/*
Add multiple descriptive attiributes with label = labels to a note with label = to
*/
func (schema *Dfm) AddSequenceDescriptive(labels string, to string) {

	for _, label := range strings.Split(labels, " ") {
		schema.AddDescriptive(label, to)
	}
}

/*
Render the diagram
*/
func (schema *Dfm) RenderDiagram() {
	output := schema.Graph.String()

	f, err := os.Create("dot.dot")
	if err != nil {
        log.Fatal(err)
    }

	defer f.Close()
	f.WriteString(output)
}