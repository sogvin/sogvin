package spec

import (
	. "github.com/gregoryv/web"
)

func NewNavigationSpec(n *Hn) *Element {
	spec := Article(
		n.H1("Navigation system"),
		Em(`Purpose; provide safe travel through space.`),

		Section(
			n.H2("Background"),

			P(`Through the navigation system people can plot a course
            or manually steer a ship.  People depend on its accuracy
            and automation to safely navigate through space.`),

			n.H3(`Plot new course`),

			P(`Standing at the bridge, the captain asks for the
			closest viable planets for some time at the beach. Selects
			the one with the nicest beaches and tells the system to
			plot the course. The plot details show that the route is
			through uncharted space. The captain selects another of
			the viable planets and tells the system to plot the course
			again. Once satisfied, he tells the system to engage.`),

			P(`The journey is estimated to five days. On the second
			day however an interference is detected in space and the
			ship adapts the course accordingly. The captain is
			notified through his personal communicator of the
			changes.`),
			//

			n.H3(`Manual control`),

			P(`Once the ship enters the planets atmosphere one of the
	 	    crew members on the bridge tells the system to let him
	 	    manually steer the ship. He wants to find a suitable spot
	 	    on the crowded beach, before letting the passengers leave
	 	    the ship.`),

			//
		),
		Section(
			n.H2("Elicited features"),
			Features(
				"voice control by multiple people",
				"show route details",
				"find destination",
			),
		),
	)
	return spec
}

func Features(v ...string) *Element {
	ul := Ul(Class("features"))
	for _, f := range v {
		ul.With(Li(f))
	}
	return ul
}
