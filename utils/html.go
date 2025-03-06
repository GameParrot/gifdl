package utils

import (
	"golang.org/x/net/html"
)

func FindHref(n *html.Node, val string) string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "aria-label" && attr.Val == val {
				for _, a := range n.Attr {
					if a.Key == "href" {
						return a.Val
					}
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if href := FindHref(c, val); href != "" {
			return href
		}
	}
	return ""
}

func Find(n *html.Node, val string, key string) *html.Node {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == key && attr.Val == val {
				return n
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if htmlK := Find(c, val, key); htmlK != nil {
			return htmlK
		}
	}
	return nil
}

func FindMeta(n *html.Node, val string, key string) string {
	if n.Type == html.ElementNode && n.DataAtom.String() == "meta" && GetAttr(n, key) == val {
		return GetAttr(n, "content")
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if htmlK := FindMeta(c, val, key); htmlK != "" {
			return htmlK
		}
	}
	return ""
}

func GetAttr(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func GetTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.DataAtom.String() == "title" && n.FirstChild != nil {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if htmlK := GetTitle(c); htmlK != "" {
			return htmlK
		}
	}
	return ""
}
