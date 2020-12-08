package spec

import (
	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/toc"
)

func NewNavigationSpec() *Element {

	nav := Nav()
	spec := Article(
		H1("Spaceship navigation system specification"),
		Em(`Purpose; provide safe travel through space.`),

		nav,
		Section(

			H2("Background"),

			P(`Through the navigation system people can plot a course
            or manually steer a ship.  People depend on its accuracy
            and automation to safely navigate through space.`),

			H3(`Plot new course`),

			P(`Standing at the bridge, the captain asks for the
			closest viable planets for some time at the beach. Selects
			the one with the nicest beaches and views and tells the
			system to plot the course. The plot details show there
			route goes through uncharted space. The captain selects
			another of the viable planets and tells the system to plot
			the course again. Once satisfied, he tells the system to
			engage.`),

			P(`The journey is estimated to five days. On the second
			day however an interference is detected in space and the
			ship adapts the course accordingly. The captain is
			notified through his personal communicator of the
			changes.`),
			//

		),
	)

	toc.MakeTOC(nav, spec, "h2", "h3")
	return spec
}