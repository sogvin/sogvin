package spec

import (
	"testing"

	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/toc"
)

func Test_specification(t *testing.T) {
	n := NewHn(2)
	nav := Nav()
	spec := NewNavigationSpec(n)
	toc.MakeTOC(nav, spec, "h2", "h3", "h4")
	NewPage(
		Html(
			Head(
				Style(Theme()),
			),
			Body(
				H1("Spaceship system specification"),
				nav,
				spec,
			),
		),
	).SaveAs("navsys.html")

}
