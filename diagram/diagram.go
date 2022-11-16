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

func (d Dfm) CreateFact(title string, attributes []string) {

	label := `<<table border="0" cellborder="1" cellspacing="0" cellpadding="20"> <tr> <td bgcolor="lightblue">`+title+`</td> </tr>`
	for _, att := range(attributes) {
		label += `<tr> <td port="port`+att+`">`+att+`</td> </tr>`
	}
	label += `</table>>`


	factAtt := map[string]string{"shape":"plain", "root":"true", "label":label}
	
	d.Graph.AddNode("G", title, factAtt)
}

func (d Dfm) AddNode(label string, attach string) {
	
	nodeAtt := map[string]string{"shape":"circle", "label":label}
	edgeAtt := map[string]string{"arrowhead":"none"}

	d.Graph.AddNode("G", label, nodeAtt)
	d.Graph.AddEdge(attach, label, true, edgeAtt)
}

func (d Dfm) AddSequenceNode(labels []string, startAttach string) {

	d.AddNode(labels[0], startAttach)
	for i := 0; i < len(labels)-1; i++ {
		d.AddNode(labels[i+1], labels[i])
	}

}

func (d Dfm) AddConvergence(label string, attach string) {

	nodeAtt := map[string]string{"shape":"circle", "label":label}
	edgeAtt := map[string]string{}

	d.Graph.AddNode("G", label, nodeAtt)
	d.Graph.AddEdge(attach, label, true, edgeAtt)
}

func (d Dfm) AddHierarchy(labels []string, from string, to string) {

	NodeAtt := map[string]string{"shape":"circle", "label":to}
	EdgeAtt := map[string]string{"arrowhead":"none"}

	d.Graph.AddNode("G", to, NodeAtt)

	for _, label := range labels {
		tmpAtt := EdgeAtt
		tmpAtt["label"] = label
		d.Graph.AddEdge(from, to, true, tmpAtt)
	}

}

func (d Dfm) AddOptional(label string, attach string) {

	nodeAtt := map[string]string{"shape":"circle", "label":label}
	edgeAtt := map[string]string{"arrowhead":"icurve"}

	d.Graph.AddNode("G", label, nodeAtt)
	d.Graph.AddEdge(attach, label, true, edgeAtt)
}

func (d Dfm) AddDescriptive(label string, to string) {

	descriptiveAtt := map[string]string{"shape":"underline"}
	edgeAtt := map[string]string{"arrowhead":"none"}

	d.Graph.AddNode("G", label, descriptiveAtt)
	d.Graph.AddEdge(to, label, true, edgeAtt)
}

func (d Dfm) AddSequenceDescriptive(labels []string, to string) {

	for _, label := range labels {
		d.AddDescriptive(label, to)
	}

}