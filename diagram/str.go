package diagram

import (
	"github.com/awalterschulze/gographviz"

	"log"
	"os"
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
Create a fact to the schema with title = name and attributes = attributes, keys = keys, table_color =  color of the table,
keys-colors the color of any key, if nil the color are the color of table, if len(keys_color) < keys return an error
*/
func (d Str) CreateFact(name string, keys []string, attributes []string, keys_color []string, table_color string) {

	if (len(keys) != len(keys_color)) {
		log.Fatal("Number of keys and number of colors for keys are different in table: "+name)
	}

	label := `<<table border="0" bgcolor="`+table_color+`" cellborder="1" cellspacing="0" cellpadding="20">`
	
	for i, key := range(keys) {
		label += `<tr> <td bgcolor="`+keys_color[i]+`" port="`+key+`"> <u>`+key+`</u> </td> </tr>`
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

/*
Add a dimension to the schema, with name = name, attributes of table = attributes, keys = keys, key = at what type of key attach 
the dimension, keys-colors the color of any key, if nil the color are the color of table, if len(keys_color) < keys return an error
*/
func (d Str) AddDimension(name string, keys[]string, attributes []string, attach string, key string, keys_color []string, table_color string) {

	d.CreateFact(name, keys, attributes, keys_color, table_color)
	d.Graph.AddPortEdge(attach, key, name, key, true, STR_edgeAtt)
}


/*
Render the diagram
*/
func (d Str) RenderDiagram() {
	output := d.Graph.String()

	f, err := os.Create("dot.dot")
	if err != nil {
        log.Fatal(err)
    }

	defer f.Close()
	f.WriteString(output)
}