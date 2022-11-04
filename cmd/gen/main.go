package main

import (
	"log"
	"os"

	"github.com/awalterschulze/gographviz"
)

func main() {
	graphAst, _ := gographviz.ParseString(`digraph G {
	}`)
	graph := gographviz.NewGraph()
	if err := gographviz.Analyse(graphAst, graph); err != nil {
		panic(err)
	}


	label := "<<table border=\"0\" cellborder=\"1\" cellspacing=\"0\" cellpadding=\"10\">" +
				"<tr> <td> sales other things just to make</td> </tr>" +
				"<tr> <td align=\"left\"> att1 </td> </tr> <tr> <td align=\"left\"> att2 </td> </tr> <tr> <td align=\"left\"> att3 </td> </tr> <tr> <td align=\"left\"> att4 </td> </tr> <tr> <td align=\"left\"> att5 </td> </tr> <tr> <td align=\"left\"> att6 </td> </tr>  </table>>"



	graph.AddSubGraph("G", "how", map[string]string{"rank": "same"})
	graph.AddNode("how", "y1", map[string]string{"shape":"box"})
	graph.AddNode("how", "y2", map[string]string{"shape":"box"})
	graph.AddNode("how", "y3", map[string]string{"shape":"box"})
	graph.AddNode("how", "y4", map[string]string{"shape":"box"})

	graph.AddNode("G", "fact", map[string]string{"shape":"plain", "label": label})

	graph.AddEdge("how", "fact", true, map[string]string{"dir":"back"})

	graph.AddSubGraph("G", "where", map[string]string{"rank": "same"})
	graph.AddNode("where", "fact", nil)

	graph.AddNode("where", "x2", map[string]string{"group": "left"})

	graph.AddEdge("x2", "fact", true, map[string]string{"dir":"back"})

	graph.AddNode("G", "x1", map[string]string{"group": "left"})
	graph.AddNode("G", "x3", map[string]string{"group": "left"})

	graph.AddEdge("x1", "x2", true, map[string]string{"style":"invis"})
	graph.AddEdge("x2", "x3", true, map[string]string{"style":"invis"})

	graph.AddEdge("fact", "x1", true, map[string]string{"dir":"back", "constraint":"false"})
	graph.AddEdge("fact", "x3", true, map[string]string{"dir":"back", "constraint":"false"})


	graph.AddNode("where", "z2", map[string]string{"group": "right"})

	graph.AddEdge("fact", "z2", true, map[string]string{"dir":"back"})

	graph.AddNode("G", "z1", map[string]string{"group": "right"})
	graph.AddNode("G", "z3", map[string]string{"group": "right"})

	graph.AddEdge("z1", "z2", true, map[string]string{"style":"invis"})
	graph.AddEdge("z2", "z3", true, map[string]string{"style":"invis"})

	graph.AddEdge("fact", "z1", true, map[string]string{"dir":"back", "constraint":"false"})
	graph.AddEdge("fact", "z3", true, map[string]string{"dir":"back", "constraint":"false"})



	output := graph.String()


	f, err := os.Create("dot.dot")
	if err != nil {
        log.Fatal(err)
    }

	defer f.Close()
	f.WriteString(output)

}