package skplst

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/awalterschulze/gographviz"
)

/* Reference graph
digraph {
  rankdir=LR
  node [shape=record,weight=4]
  edge [weight=10000]
  nodesep=0

  X [label="<f0>•|<f1>•|<f2>•|<f3>•|<f4>Head"]
  A [label="<f3>•|<f4>4"]
  B [label="<f1>•|<f2>•|<f3>•|<f4>8"]
  C [label="<f3>•|<f4>15"]
  D [label="<f0>•|<f1>•|<f2>•|<f3>•|<f4>16"]
  E [label="<f2>•|<f3>•|<f4>23"]
  F [label="<f2>•|<f3>•|<f4>42"]
  Y [label="<f0>•|<f1>•|<f2>•|<f3>•|<f4>Tail"]

  X:f0 -> D:f0
  X:f1 -> B:f1
  X:f2 -> B:f2
  X:f3 -> A:f3
  X:f4 -> A:f4

  A:f3 -> B:f3
  A:f4 -> B:f4

  B:f1 -> D:f1
  B:f2 -> D:f2
  B:f3 -> C:f3
  B:f4 -> C:f4

  C:f3 -> D:f3
  C:f4 -> D:f4

  D:f0 -> Y:f0
  D:f1 -> Y:f1
  D:f2 -> E:f2
  D:f3 -> E:f3
  D:f4 -> E:f4

  E:f3 -> F:f3
  E:f4 -> F:f4

  F:f2 -> Y:f2
  F:f3 -> Y:f3
  F:f4 -> Y:f4
}
*/

func Dot(list *SkipList) *gographviz.Graph {
	graph := gographviz.NewGraph()
	graph.AddAttr("", "rankdir", "LR")
	graph.AddAttr("", "nodesep", "0")

	graph.AddNode("", "node", map[string]string{
		"shape":  "record",
		"weight": "4",
	})
	graph.AddNode("", "edge", map[string]string{
		"weight": "10000",
	})

	// head & tail
	graph.AddNode("", "head", nil)
	// the first port is node name head/tail/<key>
	AddRecordPort(graph.Nodes.Lookup["head"], fmt.Sprintf("f%d", 0), "head")
	for i := 0; i < len(list.Forward); i++ {
		AddRecordPort(graph.Nodes.Lookup["head"], fmt.Sprintf("f%d", i+1), "•")
	}

	graph.AddNode("", "tail", nil)
	AddRecordPort(graph.Nodes.Lookup["tail"], fmt.Sprintf("f%d", 0), "tail")
	for i := 0; i < len(list.Forward); i++ {
		AddRecordPort(graph.Nodes.Lookup["tail"], fmt.Sprintf("f%d", i+1), "•")
	}
	for i := 0; i < 10; i++ {
		curNode := graph.Nodes.Lookup[strconv.Itoa(i)]
		listNode := nil
		AddRecordPort(curNode, fmt.Sprintf("f%d", 0), strconv.Itoa(i))
		for i := 0; i < len(listNode.Forward); i++ {
			AddRecordPort(curNode, fmt.Sprintf("f%d", i+1), "•")
		}
	}
	// add edges
	for {
	}

	return graph
}

// see https://graphviz.org/doc/info/shapes.html#record for the grammar of parsing port record
var pnamelabelReg = regexp.MustCompile(`<(.*)>(.*)`)

// AddRecordPort add a port to record tpe node
func AddRecordPort(node *gographviz.Node, name, label string) {
	var port string
	if name != "" {
		port = fmt.Sprintf("<%s>%s", name, label)
	} else {
		port = fmt.Sprintf("%s", label)
	}

	labels := node.Attrs["label"]
	labelList := strings.Split(labels, "|")
	replaced := false

	for i := range labelList {
		fields := pnamelabelReg.FindStringSubmatch(labelList[i])
		if len(fields) == 3 { // port is of the form  <name>label
			if fields[1] == name { // field already exists, update label
				labelList[i] = port
				replaced = true
			}
		}
	}
	if !replaced { // append new label
		labelList = append(labelList, port)
	}

	// final: join and replace label
	node.Attrs["label"] = strings.Join(labelList, "|")

}
