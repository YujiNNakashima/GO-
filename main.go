package main

import (
	"fmt"
	queryselector "softwaredesign/query_selector"
)

func main() {

	root := &queryselector.Node{
		Type: "tag",
		Name: "html",
		Attribs: map[string]string{
			"id":    "root",
			"class": "main",
		},
		Children: []*queryselector.Node{
			{
				Type: "tag",
				Name: "body",
				Attribs: map[string]string{
					"id":    "body",
					"class": "content",
				},
				Children: []*queryselector.Node{
					{
						Type:    "tag",
						Name:    "div",
						Attribs: map[string]string{"id": "div1"},
					},
					{
						Type:    "tag",
						Name:    "div",
						Attribs: map[string]string{"id": "div2"},
					},
					{
						Type: "tag",
						Name: "span",
					},
				},
			},
		},
	}

	selectedNode := queryselector.Select(root, "body div#div2")
	fmt.Println(selectedNode)
}
