package spec

import (
	"testing"
)

func Test_specification(t *testing.T) {
	page := NewPage(
		Html(
			Head(
				Style(Theme()),
			),
			Body(
				NewNavigationSpec(),
			),
		),
	).SaveAs("navsys.html")

}
