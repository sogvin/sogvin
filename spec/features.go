package spec

import . "github.com/gregoryv/web"

func NewElicitedFeatures(n *Hn) *Element {
	return Article(
		n.H1("Elicited features"),

		P(`These features have been elicitated from the navigation story`),

		n.H2("voice control"),
		n.H2("show route details"),
		n.H2("find destination"),
	)
}
