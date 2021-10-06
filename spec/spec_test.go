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
				Meta(Charset("utf-8")),
				Style(theme()),
				Script(
					// to prevent Firefox FOUC, this must be here
					// https://stackoverflow.com/questions/21147149
					"let FF_FOUC_FIX;",
				),
			),
			body,
		),
	).SaveAs("docs/index.html")
}
