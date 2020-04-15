package util

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func ParseNodeToStringSliceUseHtmlquery(root *html.Node, xpath string) []string {
	var rt []string
	ne := htmlquery.Find(root, xpath)
	for _, v := range ne {
		rt = append(rt, htmlquery.InnerText(v))
	}
	return rt
}

func ParseNodeToStringUseHtmlquery(root *html.Node, xpath string) string {
	rtBuilder := ParseNodeToStringSliceUseHtmlquery(root, xpath)
	if len(rtBuilder) == 0 {
		return ""
	}
	return rtBuilder[0]
}
