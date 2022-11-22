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
	STR_factAtt =  map[string]string{"shape":"plaintext", "root":"true"}
	STR_edgeAtt = map[string]string{"arrowhead":"none"}

)

/*
Create a new object *Str to use for creating the star schema
*/
func NewSTR() *Str {

	graphAst, _ := gographviz.ParseString(`digraph G { 
		layout="circo"
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
func (d Str) CreateFact(name string, keys []string, attributes []string) {
	
	label := `<<table border="0" cellborder="1" cellspacing="0" cellpadding="20">`
	
	for _, key := range(keys) {
		label += `<tr> <td port="`+key+`"> <u>`+key+`</u> </td> </tr>`
	}
	
	for _, att := range(attributes) {
		label += `<tr> <td port="`+att+`">`+att+`</td> </tr>`
	}

	label += `</table>>`

	fact_att := STR_factAtt
	fact_att["xlabel"] = `<<font point-size="20">`+name+`</font>>`
	fact_att["label"] = label

	d.Graph.AddNode("G", name, fact_att)
}

func (d Str) AddDimension(name string, keys[]string, attributes []string, attach string, key string) {

	d.CreateFact(name, keys, attributes)

	d.Graph.AddPortEdge(attach, key, name, key, true, STR_edgeAtt)
}