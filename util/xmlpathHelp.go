package util

import (
	"gopkg.in/xmlpath.v1"
	"strconv"
)

func ParseNodeToStringSlice(xpath string, node *xmlpath.Node) []string {
	var rt []string
	xPath := xmlpath.MustCompile(xpath)
	it := xPath.Iter(node)
	for it.Next() {
		rt = append(rt, it.Node().String())
	}
	return rt
}

func ParseNodeToIntSlice(xpath string, node *xmlpath.Node) []int {
	rtBuilder := ParseNodeToStringSlice(xpath, node)
	var rt []int
	for _, v := range rtBuilder {
		e, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		rt = append(rt, e)
	}
	return rt
}

func ParseNodeToString(xpath string, node *xmlpath.Node) string {
	rt := ParseNodeToStringSlice(xpath, node)
	return rt[0]
}
