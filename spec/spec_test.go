package spec

import (
	"strings"
	"testing"

	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/toc"
)

func Test_specification(t *testing.T) {
	n := NewHn(2)
	nav := Nav()
	spec := NewNavigationSpec(n)
	features := NewElicitedFeatures(n)

	LinkAll(features, "navigation story", "navigationsystem")

	body := Body(
		H1("Spaceship system specification"),
		nav,
		spec,
		features,
	)

	toc.MakeTOC(nav, body, "h2", "h3", "h4")
	NewPage(
		Html(
			Head(
				Style(Theme()),
			),
			body,
		),
	).SaveAs("navsys.html")

}

func LinkAll(el *Element, txt string, toId string) {

	for cIndex, c := range el.Children {
		switch c := c.(type) {
		case string:
			i := strings.Index(c, txt)
			if i >= 0 {
				before := c[:i]
				lnk := A(Href("#"+toId), txt)
				after := c[i+len(txt):]
				el.Children[cIndex] = Span(before, lnk, after)
			}
		case *Element:
			LinkAll(c, txt, toId) // recursive
		}
	}
}
