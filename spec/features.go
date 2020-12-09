package spec

import . "github.com/gregoryv/web"

func NewElicitedFeatures(n *Hn) *Element {
	return Article(
		n.H1("Elicited features"),
		Features(
			"voice control by multiple people",
			"show route details",
			"find destination",
		),
	)
}
