package diagram

import (
	"github.com/awalterschulze/gographviz"

	"log"
	"os"
	"strings"
)


var colors = []string{"crimson", "darkcyan", "green"}

/*
Struct for STR schema
*/
type Str struct {
	Graph *gographviz.Graph	
	colors map[string]string
	index_color uint8
}

type Dimension struct {
	name string
	keys []string
	attributes []string
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
		make(map[string]string),
		0,
	}
}

/*
Create dimension
*/
func (d *Str) CreateDimension(name string, keys string, attributes string) *Dimension {
	
	t_keys := strings.Split(keys, " ")
	t_attributes := strings.Split(attributes, " ")

	
	dimension := &Dimension{
		name,
		t_keys,
		t_attributes,
	}

	d.RenderDimension(dimension)

	return dimension
}

/*
Create a fact to the schema with title = name and attributes = attributes, keys = keys, table_color =  color of the table,
keys-colors the color of any key, if nil the color are the color of table, if len(keys_color) < keys return an error
*/
func (d *Str) RenderDimension(dim *Dimension) {

	color := d.colors[dim.keys[0]]
	if color == "" {
		color = "white"
	}

	label := `<<table border="0" bgcolor="`+color+`" cellborder="1" cellspacing="0" cellpadding="20">`
	
	for _, key := range(dim.keys) {
		color := d.colors[key]
		if color == "" {
			color = colors[d.index_color]
			d.colors[key] = color
			d.index_color++
		}

		label += `<tr> <td bgcolor="`+color+`" port="`+key+`"> <u>`+key+`</u> </td> </tr>`
	}
	
		for _, att := range(dim.attributes) {
			if att != "" {
				label += `<tr> <td port="`+att+`">`+att+`</td> </tr>`
			}
		}

	label += `</table>>`

	fact_att := STR_factAtt
	fact_att["xlabel"] = `<<font point-size="20">`+dim.name+`</font>>`
	fact_att["label"] = label

	d.Graph.AddNode("G", dim.name, fact_att)
}

/*
Add a dimension to the schema, with name = name, attributes of table = attributes, keys = keys, key = at what type of key attach 
the dimension, keys-colors the color of any key, if nil the color are the color of table, if len(keys_color) < keys return an error
*/
func (d *Str) JoinDimension(d1 *Dimension, d2 *Dimension, portKey string) {
	d.Graph.AddPortEdge(d1.name, portKey, d2.name, portKey, true, STR_edgeAtt)
}


/*
Render the diagram
*/
func (d *Str) RenderDiagram() {
	output := d.Graph.String()

	f, err := os.Create("dot.dot")
	if err != nil {
        log.Fatal(err)
    }

	defer f.Close()
	f.WriteString(output)
}