package htmldom

import "golang.org/x/net/html"

// inspired by https://siongui.github.io/2016/04/15/go-getElementById-via-net-html-package/

func GetFirstElementByClass(node *html.Node, id string) *html.Node {
	return traverse(node, id)
}

func traverse(node *html.Node, id string) *html.Node {
	if checkID(node, id) {
		return node
	}

	for current := node.FirstChild; current != nil; current = current.NextSibling {
		result := traverse(current, id)
		if result != nil {
			return result
		}
	}

	return nil
}

func checkID(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		s, ok := GetAttribute(n, "class")
		if ok && s == id {
			return true
		}
	}
	return false
}

func GetAttribute(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}
