package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link struct
type Link struct {
	Href string
	Text string
}

// Parse func
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	fmt.Println("node...", n)
	ret.Text = text(n)
	return ret
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c) + " "
	}
	return strings.Join(strings.Fields(ret), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	fmt.Println("doc...", n)
	if n.Type == html.ElementNode && n.Data == "a" {
		fmt.Printf("true??\n")
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		fmt.Printf("for loop ret ==\n%+v", ret)
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
