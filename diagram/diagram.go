package diagram

import (
	"github.com/awalterschulze/gographviz"
)

type Dfm struct {
	Graph *gographviz.Graph	 //graph
}

/*
Attributes for modify nodes in the graph
*/
var (
	nodeAtt = map[string]string{"shape":"circle"}
	edgeAtt = map[string]string{"arrowhead":"none"}
	factAtt =  map[string]string{"shape":"plain", "root":"true"}
	descriptiveAtt = map[string]string{"shape":"underline"}
	optionalAtt = map[string]string{"arrowhead":"icurve"}
	hierarchyAtt = map[string]string{"arrowhead":"none"}
)

/*
NewDfm() return a Dfm struct with gragh
*/
func NewDFM() *Dfm {

	graphAst, _ := gographviz.ParseString(`digraph G { 
		layout="circo"
		overlap=false
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
CreateFact(title string, attributes []string) add a fact to a node with title = title and attributes = attiributes
*/
func (d Dfm) CreateFact(title string, attributes []string) {

	label := `<<table border="0" cellborder="1" cellspacing="0" cellpadding="20"> <tr> <td bgcolor="lightblue">`+title+`</td> </tr>`
	for _, att := range(attributes) {
		label += `<tr> <td port="port`+att+`">`+att+`</td> </tr>`
	}
	label += `</table>>`

	fact_att := factAtt
	fact_att["label"] = label

	d.Graph.AddNode("G", title, fact_att)
}

/*
AddNode(label string, attach string) add a node with label = label to a node with label = attach
AddSequenceNode(labels[] string, startAttach string) add multiple nodes with label = labels, starting from node with label = startAttach to the node with label = labels[len(labels)]
*/
func (d Dfm) AddNode(label string, attach string) {
	
	node_att := nodeAtt
	node_att["label"] = label

	d.Graph.AddNode("G", label, node_att)
	d.Graph.AddEdge(attach, label, true, edgeAtt)
}

func (d Dfm) AddSequenceNode(labels []string, startAttach string) {

	d.AddNode(labels[0], startAttach)
	for i := 0; i < len(labels)-1; i++ {
		d.AddNode(labels[i+1], labels[i])
	}

}

/*
addConvergence(label string, attach string) add a convergence node with label = label to a node with label = attach
*/
func (d Dfm) AddConvergence(label string, attach string) {
	node_att := nodeAtt
	node_att["label"] = label

	d.Graph.AddNode("G", label, node_att)
	d.Graph.AddEdge(attach, label, true, nil)
}

/*
AddHierarchy(labels []string, from string, to string) add a hierachy with 2 or more node with label = labels starting from node with label = to
*/
func (d Dfm) AddHierarchy(labels []string, from string, to string) {

	node_att := nodeAtt
	node_att["label"] = to

	d.Graph.AddNode("G", to, node_att)

	for _, label := range labels {
		tmpAtt := hierarchyAtt
		tmpAtt["label"] = label
		d.Graph.AddEdge(from, to, true, tmpAtt)
	}

}

/*
AddOptional(label string, attach string) add an optional node with label = label starting from a node with label = attach
*/
func (d Dfm) AddOptional(label string, attach string) {

	node_att := nodeAtt
	node_att["label"] = label

	d.Graph.AddNode("G", label, node_att)
	d.Graph.AddEdge(attach, label, true, optionalAtt)
}


/*
AddDescriptive(label string, to string) add a new descriptive attribute with the label = label to a node with label = to
AddSequenceDescriptive(labels []string, to string) add multiple descriptive attiributes with label = labels to a note with label = to
*/
func (d Dfm) AddDescriptive(label string, to string) {

	d.Graph.AddNode("G", label, descriptiveAtt)
	d.Graph.AddEdge(to, label, true, edgeAtt)
}

func (d Dfm) AddSequenceDescriptive(labels []string, to string) {

	for _, label := range labels {
		d.AddDescriptive(label, to)
	}

}