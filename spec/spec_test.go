package spec

import (
	"testing"

	. "github.com/gregoryv/web"
)

func Test_index(t *testing.T) {
	n := NewHn(1)
	body := Body(
		NewExploreRequirementsEngineering(n),
		NewBeachStory(n),
	)
	NewPage(
		Html(
			Head(
				Style(Theme()),
			),
			body,
		),
	).SaveAs("docs/index.html")
}
