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
						Attribs: map[string]string{"id": "my-id-1"},
					},
					{
						Type:    "tag",
						Name:    "div",
						Attribs: map[string]string{"id": "my-id-2"},
					},
					{
						Type: "tag",
						Name: "span",
					},
				},
			},
		},
	}

	selectedNode := queryselector.Select(root, "body div#my-id-2")
	fmt.Println(selectedNode)
}
