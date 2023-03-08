package diagram

import (
	"github.com/awalterschulze/gographviz"

	"log"
	"os"
	"strings"
)


var colors = []string{"crimson", "darkcyan", "green", "brown", "beige", "gold", "indigo", "lime", "magenta", "navy", "yellow", "cornflowerblue"}

/*
Struct Str defines the star schema and includes a pointer to 
a graph, a map of colors, and an index for the color array.
*/
type Str struct {
	Graph *gographviz.Graph	
	colors map[string]string
	index_color int
}


/*
Struct Dimension defines a dimension and includes a name, keys, 
and attributes.
*/
type Dimension struct {
	name string
	keys []string
	attributes []string
}

/*
Attributes are maps of attributes for modifying nodes and edges 
in the star schema.
*/
var (
	STR_factAtt =  map[string]string{"shape":"plaintext", "root":"true"}
	STR_edgeAtt = map[string]string{"arrowhead":"none"}
)

/*
Function NewSTR creates a new Str object for creating the star schema.
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
Function CreateDimension creates a new Dimension object and renders 
the dimension in the schema.
*/
func (schema *Str) CreateDimension(name string, keys string, attributes string) *Dimension {
	
	t_keys := strings.Split(keys, " ")
	t_attributes := strings.Split(attributes, " ")

	
	dimension := &Dimension {
		name,
		t_keys,
		t_attributes,
	}

	schema.RenderDimension(dimension)

	return dimension
}

/*
Function RenderDimension renders a given dimension in the schema.
*/
func (schema *Str) RenderDimension(dim *Dimension) {

	color := schema.colors[dim.keys[0]]
	if color == "" {
		color = "white"
	}

	label := `<<table border="0" bgcolor="`+color+`" cellborder="1" cellspacing="0" cellpadding="20">`
	
	for _, key := range(dim.keys) {
		color := schema.colors[key]
		if color == "" {
			color = colors[schema.index_color % len(colors)]
			schema.colors[key] = color
			schema.index_color++
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

	schema.Graph.AddNode("G", dim.name, fact_att)
}

/*
Function JoinDimension joins two dimensions by a portkey.
*/
func (schema *Str) JoinDimension(d1 *Dimension, d2 *Dimension, portKey string) {
	schema.Graph.AddPortEdge(d1.name, portKey, d2.name, portKey, true, STR_edgeAtt)
}


/*
Function RenderDiagram renders the entire schema to a dot file.
*/
func (schema *Str) RenderDiagram() {
	output := schema.Graph.String()

	f, err := os.Create("dot.dot")
	if err != nil {
        log.Fatal(err)
    }

	defer f.Close()
	f.WriteString(output)
}