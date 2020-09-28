package sogvin

import (
	"fmt"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/shape"
	. "github.com/gregoryv/web"
)

var componentsDiagram = Article(
	H1("Components diagram"),
	P(
		`The components diagram shows services and processes related
		to one another. Good for system overviews and microservice
		architectures.`,
	),
	Span(Class("Center"),
		newOverviewDiagram(),
	),
	P(
		`Use lines between components unless you are conveying
		signaling direction between them.`,
	),
	H2("General diagram tips"),
	P(
		`In general diagrams should be kept simple, 5-9 items is a
		good rule (same as complexity in code).`,
	),
	Ul(
		Li("Highlight important components"),
		Li("Use arrows when direction is important otherwise plain lines"),
		Li(
			`Reflect on cognitive placement, ie. cloud components are
			above others`,
		),
		Li(
			`Use white as emphasizing color, works in both grayscale
			 and colored diagrams`,
		),
		Li("Stick to one color scheme"),
	),
	colorSchemeDiagram(),
)

func newOverviewDiagram() *Element {
	draw.ClassAttributes["external"] = `stroke="#d3d3d3" fill="#e2e2e2"`
	var (
		d        = design.NewClassDiagram()
		serviced = shape.NewComponent("serviced")
		ng       = shape.NewComponent("nginx")
		inet     = shape.NewCircle(40)
		client   = shape.NewComponent("client")
		db       = shape.NewDatabase("postgres")
		cloud    = shape.NewLabel("internet")
	)
	shape.SetClass("external", ng, client, inet, db)

	d.Place(serviced).At(130, 20)
	d.Place(db).RightOf(serviced)
	d.HAlignCenter(serviced, db)
	d.Place(ng).Below(serviced, 40)
	d.VAlignCenter(serviced, ng)
	d.Place(inet).LeftOf(ng, 70)
	d.Place(cloud).RightOf(inet)
	d.VAlignCenter(inet, cloud)

	d.HAlignCenter(ng, inet)
	d.Place(client).Below(inet, 40)
	d.VAlignCenter(inet, client)

	lineBetween := func(a, b shape.Shape) {
		d.Link(a, b).Head = nil
	}
	lineBetween(serviced, db)
	lineBetween(ng, serviced)
	lineBetween(ng, inet)
	lineBetween(inet, client)

	src := "img/overview.svg"
	d.SaveAs("htdocs/" + src)
	return Img(Src(src))
}

func colorSchemeDiagram() *Element {
	var (
		d      = design.NewDiagram()
		colors = []string{
			"#ffffff",
			"#e2e2e2",
			"#ffffcc",
			"#ffcc99",
			"#ff9999",
			"#ccff99",
			"#99e6ff",
		}
	)
	var last shape.Shape
	for i, color := range colors {
		class := fmt.Sprintf("circle%v", i)
		v := fmt.Sprintf(`stroke="#d3d3d3" stroke-width="1" fill="%s"`, color)
		draw.ClassAttributes[class] = v
		c := shape.NewCircle(30)
		c.SetClass(class)
		l := shape.NewLabel(color)
		if last == nil {
			d.Place(c).At(20, 20)
		} else {
			d.Place(c).RightOf(last)
		}
		last = c
		d.Place(l).RightOf(c)
		d.VAlignCenter(c, l)
		shape.Move(l, 0, 15)
	}

	src := "img/color_scheme.svg"
	d.SaveAs("htdocs/" + src)
	return Img(Src(src))
}
