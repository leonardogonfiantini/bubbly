package diagram

import (
	"github.com/awalterschulze/gographviz"
)

/*
Struct for STR schema
*/
type Str struct {
	Graph *gographviz.Graph	 //graph
}

/*
Attributes for modify nodes in the STR schema
*/
var (
	STR_factAtt =  map[string]string{"shape":"plain", "root":"true"}
)

/*
Create a new object *Str to use for creating the star schema
*/
func NewSTR() *Str {

	graphAst, _ := gographviz.ParseString(`digraph G { 
	}`)
	graph := gographviz.NewGraph()
	if err := gographviz.Analyse(graphAst, graph); err != nil {
		panic(err)
	}

	return &Str{
		graph,
	}
}

/*
Create a fact to the schema with title = title and attributes = attributes
*/
func (d Str) CreateFactTable(title string, attributes []string) {
	
	label := `<<table border="0" cellborder="1" cellspacing="0" cellpadding="20"> <tr> <td bgcolor="lightblue">`+title+`</td> </tr>`
	for _, att := range(attributes) {
		label += `<tr> <td port="port.`+att+`">`+att+`</td> </tr>`
	}
	label += `</table>>`

	fact_att := STR_factAtt
	fact_att["label"] = label

	d.Graph.AddNode("G", title, fact_att)
}