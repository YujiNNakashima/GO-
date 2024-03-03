package queryselector

import (
	"fmt"
	"strings"
)

type Node struct {
	Type     string
	Name     string
	Attribs  map[string]string
	Children []*Node
}

func Select(root *Node, selector string) *Node {
	selectors := strings.Fields(selector)
	return FirstMatch(root, selectors)
}

func FirstMatch(node *Node, selectors []string) *Node {
	if len(selectors) == 0 {
		panic("Require selector(s)")
	}

	if node.Type != "tag" {
		return nil
	}

	if MatchHere(node, selectors[0]) {
		if len(selectors) == 1 {
			return node
		}
		return FirstChildMatch(node, selectors[1:])
	}

	return FirstChildMatch(node, selectors)
}

func FirstChildMatch(node *Node, selectors []string) *Node {
	if node.Type != "tag" {
		panic(fmt.Sprintf("Should only try to match first child of tags, not %s", node.Type))
	}

	for _, child := range node.Children {
		match := FirstMatch(child, selectors)
		if match != nil {
			return match
		}
	}

	return nil
}

func MatchHere(node *Node, selector string) bool {
	var name, id, cls string
	if strings.Contains(selector, "#") {
		parts := strings.SplitN(selector, "#", 2)
		name, id = parts[0], parts[1]
	} else if strings.Contains(selector, ".") {
		parts := strings.SplitN(selector, ".", 2)
		name, cls = parts[0], parts[1]
	} else {
		name = selector
	}

	return node.Name == name &&
		(id == "" || node.Attribs["id"] == id) &&
		(cls == "" || strings.Contains(node.Attribs["class"], cls))
}
