package diagram

import (
	"github.com/awalterschulze/gographviz"
)



var nodeAtt = map[string]string{"width":"0.5", "label":"\"\""}
var edgeAtt = map[string]string{}


type Dfm struct {
	Graph *gographviz.Graph
}

func NewDFM() *Dfm {

	graphAst, _ := gographviz.ParseString(`digraph G { 
		layout=circo
	}`)
	graph := gographviz.NewGraph()
	if err := gographviz.Analyse(graphAst, graph); err != nil {
		panic(err)
	}

	return &Dfm{
		graph,
	}
}

func (d Dfm) AddNode(label string, attach string) {
	
	d.Graph.AddNode("G", label, nodeAtt)
	
	edgeMap := edgeAtt
	edgeMap["label"] = label
	d.Graph.AddEdge(attach, label, true, edgeMap)

}