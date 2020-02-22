package sogvin

import . "github.com/gregoryv/web"

func toc(book *Book) {
	toc := Article(Class("toc"),
		H1("Software Engineering"),
		Img(Src("img/office.jpg")),
		P("Notes by ", myname),

		H2("Design"),
		Ul(
			book.AddPage("Design", purposeOfFuncMain),
			book.AddPage("Design", nexusPattern),
			book.AddPage("Design", gracefulServerShutdown),
			book.AddPage("Design", componentsDiagram),
		),
		H3("Go packages"),
		Ul(
			gregoryv("draw", "software engineering diagrams"),
			gregoryv("web", "html generation"),
		),

		H2("Test"),
		Ul(
			book.AddPage("Test", inlineTestHelpers),
			book.AddPage("Test", alternateDesign),
		),
		H3("Go packages"),
		Ul(
			gregoryv("golden", "simplify use of golden files"),
			gregoryv("qual", "quality constraints"),
			gregoryv("ex", "indented JSON or redirect handler response to stdout"),
			gregoryv("uncover", "paths that need more testing"),
		),

		H2("Build"),
		Ul(
			book.AddPage("Build", embedVersionAndRevision),
		),
		H3("Go packages"),
		Ul(
			gregoryv("stamp", "build information code generator"),
			gregoryv("find", "files by name or content"),
		),
	)
	index := newPage("index.html", PageHeader(""), toc, Footer())
	book.pages = append(book.pages, index)
}
