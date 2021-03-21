package x_net_html

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"strings"
	"testing"
)

func TestNode(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(`<html><head></head><body><a href="foo">Foo</a><h1>hello world</h1></body></html>`))
	if err != nil {
		log.Fatal(err)
	}
	var fn func(*html.Node)
	fn = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		if n.Type == html.ElementNode && n.Data == "h1" {
			fmt.Println(n.FirstChild.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			fn(c)
		}
	}
	fn(doc)
}
