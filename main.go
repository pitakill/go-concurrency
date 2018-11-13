package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	urlFlag := flag.String("url", "https://google.com/", "URL to crawl")

	flag.Parse()

	data, err := getData(*urlFlag)
	if err != nil {
		log.Fatal(err)
	}

	html, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	iterator(html)
}

func iterator(n *html.Node) {
	if n.Type == html.ElementNode && n.DataAtom == atom.A {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(a.Val)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		iterator(c)
	}
}
