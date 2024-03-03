package queryselector

import (
	"testing"
)

func TestSelect(t *testing.T) {
	root := &Node{
		Type: "tag",
		Name: "html",
		Attribs: map[string]string{
			"id":    "root",
			"class": "main",
		},
		Children: []*Node{
			{
				Type: "tag",
				Name: "body",
				Attribs: map[string]string{
					"id":    "body",
					"class": "content",
				},
				Children: []*Node{
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

	selectedNode := Select(root, "body div#my-id-2")
	if selectedNode == nil || selectedNode.Name != "div" || selectedNode.Attribs["id"] != "my-id-2" {
		t.Errorf("Select(root, \"body div#my-id-2\") returned unexpected result")
	}

	selectedNode = Select(root, "body span#my-id-3")
	if selectedNode != nil {
		t.Errorf("Select(root, \"body span#my-id-3\") returned non-nil result")
	}
}
