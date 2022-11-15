package diagram

import (
	"github.com/awalterschulze/gographviz"
)

type Dfm struct {
	Graph *gographviz.Graph
}

func NewDFM() *Dfm {

	graphAst, _ := gographviz.ParseString(`digraph G { 
		layout="circo"
	}`)
	graph := gographviz.NewGraph()
	if err := gographviz.Analyse(graphAst, graph); err != nil {
		panic(err)
	}

	return &Dfm{
		graph,
	}
}

func (d Dfm) CreateFact(title string, attributes []string) {

	label := "<<table border=\"0\" cellborder=\"0\" cellspacing=\"0\" cellpadding=\"10\"> <tr> <td>"+title+"</td> </tr>"
	for _, att := range(attributes) {
		label += "<tr> <td>"+att+"</td> </tr>"
	}
	label += "</table>>"


	factAtt := map[string]string{"shape":"box", "root":"true", "label":label }
	
	d.Graph.AddNode("G", "fact", factAtt)

}

func (d Dfm) AddNode(label string, attach string) {
	
	nodeAtt := map[string]string{"shape":"circle", "label":"\"\"", "xlabel":label}
	edgeAtt := map[string]string{"arrowhead":"none"}

	d.Graph.AddNode("G", label, nodeAtt)
	d.Graph.AddEdge(attach, label, true, edgeAtt)

}

func (d Dfm) AddConvergence(label string, attach string) {
	nodeAtt := map[string]string{"shape":"circle", "label":"\"\"", "xlabel":label}
	edgeAtt := map[string]string{}

	d.Graph.AddNode("G", label, nodeAtt)
	d.Graph.AddEdge(attach, label, true, edgeAtt)
}

func (d Dfm) AddOptional(label string, attach string) {
	nodeAtt := map[string]string{"shape":"circle", "label":"\"\"", "xlabel":label}
	edgeAtt := map[string]string{"arrowhead":"icurve"}

	d.Graph.AddNode("G", label, nodeAtt)
	d.Graph.AddEdge(attach, label, true, edgeAtt)
}

func (d Dfm) AddDescriptive(to string, label string) {

	descriptiveAtt := map[string]string{"shape":"underline"}
	edgeAtt := map[string]string{"arrowhead":"none"}

	d.Graph.AddNode("G", label, descriptiveAtt)
	d.Graph.AddEdge(to, label, true, edgeAtt)


}