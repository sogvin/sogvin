package spec

import (
	. "github.com/gregoryv/web"
)

func NewNavigationSystem(n *Hn) *Element {

	return Article(
		n.H1("Navigation system"),

		Em(`Purpose; provide safe travel through space.`),

		P(`Through the navigation system people can plot a course or
        manually steer a ship.  People depend on its accuracy and
        automation to safely navigate through space.`),

		//
	)
}
