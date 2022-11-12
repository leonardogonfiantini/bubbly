package main

import (
	"log"
	"os"

	"github.com/leonardogonfiantini/bubbly/diagram"
)

func main() {
	
	label := "<<table border=\"0\" cellborder=\"1\" cellspacing=\"0\" cellpadding=\"10\">" +
				"<tr> <td port=\"1\"> sales other things just to make</td> </tr>" +
				"<tr> <td align=\"left\" port=\"2\"> att1 </td> </tr> <tr> <td align=\"left\" port=\"3\"> att2 </td> </tr> <tr> <td align=\"left\"> att3 </td> </tr> <tr> <td align=\"left\"> att4 </td> </tr> <tr> <td align=\"left\"> att5 </td> </tr> <tr> <td align=\"left\"> att6 </td> </tr>  </table>>"


	d := diagram.NewDFM()
	graph := d.Graph

	graph.AddNode("G", "fact", map[string]string{"shape":"box", "root":"true", "label":label })

	d.AddNode("ciao", "fact")




	output := graph.String()


	f, err := os.Create("dot.dot")
	if err != nil {
        log.Fatal(err)
    }

	defer f.Close()
	f.WriteString(output)

}