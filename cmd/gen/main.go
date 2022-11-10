package main

import (
	"log"
	"os"

	"github.com/awalterschulze/gographviz"
)

func main() {
	graphAst, _ := gographviz.ParseString(`digraph G { 
		layout=circo
	}`)
	graph := gographviz.NewGraph()
	if err := gographviz.Analyse(graphAst, graph); err != nil {
		panic(err)
	}
	label := "<<table border=\"0\" cellborder=\"1\" cellspacing=\"0\" cellpadding=\"10\">" +
				"<tr> <td port=\"1\"> sales other things just to make</td> </tr>" +
				"<tr> <td align=\"left\" port=\"2\"> att1 </td> </tr> <tr> <td align=\"left\" port=\"3\"> att2 </td> </tr> <tr> <td align=\"left\"> att3 </td> </tr> <tr> <td align=\"left\"> att4 </td> </tr> <tr> <td align=\"left\"> att5 </td> </tr> <tr> <td align=\"left\"> att6 </td> </tr>  </table>>"


	graph.AddNode("G", "fact", map[string]string{"shape":"box", "root":"true", "label":label })

	graph.AddSubGraph("G", "dom1", map[string]string{"rankdir":"LR"})
	graph.AddNode("dom1", "a", nil)
	graph.AddNode("dom1", "b", nil)
	graph.AddNode("dom1", "c", nil)
	graph.AddNode("dom1", "e", nil)
	graph.AddNode("dom1", "x", nil)
	graph.AddNode("dom1", "y", nil)
	graph.AddNode("dom1", "w", nil)
	graph.AddNode("dom1", "z", nil)

	graph.AddEdge("a", "b", true, nil)
	graph.AddEdge("a", "c", true, nil)
	graph.AddEdge("c", "e", true, nil)
	graph.AddEdge("b", "e", true, nil)

	graph.AddEdge("e", "z", true, nil)
	graph.AddEdge("z", "w", true, nil)
	graph.AddEdge("w", "x", true, nil)
	graph.AddEdge("w", "y", true, nil)
	graph.AddEdge("y", "x", true, nil)


	graph.AddEdge("fact", "a", true, nil)
	graph.AddEdge("fact", "a", true, nil)


	graph.AddSubGraph("G", "dom2", nil)
	graph.AddNode("dom2", "f", nil)
	graph.AddNode("dom2", "g", nil)
	graph.AddNode("dom2", "h", nil)
	graph.AddNode("dom2", "i", nil)

	graph.AddNode("G", "k", nil)
	graph.AddNode("G", "kk", nil)
	graph.AddNode("G", "kkk", nil)
	graph.AddNode("G", "kkkk", nil)

	graph.AddEdge("fact", "k", true, nil)
	graph.AddEdge("k", "kk", true, nil)
	graph.AddEdge("kk", "kkk", true, nil)
	graph.AddEdge("kkk", "kkkk", true, nil)


	graph.AddEdge("f", "g", true, nil)
	graph.AddEdge("h", "i", true, nil)
	graph.AddEdge("g", "i", true, nil)

	graph.AddEdge("fact", "f", true, nil)
	graph.AddEdge("fact", "h", true, nil)





	output := graph.String()


	f, err := os.Create("dot.dot")
	if err != nil {
        log.Fatal(err)
    }

	defer f.Close()
	f.WriteString(output)

}