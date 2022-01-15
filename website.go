package sogvin

import (
	"os"
	"path/filepath"

	. "github.com/gregoryv/web"
)

func NewWebsite() *Website {
	site := Website{
		Title:  "Software Engineering",
		Author: "Gregory Vin&ccaron;i&cacute;",
	}
	site.AddThemes(a4(), theme())

	toc := Article(Class("toc"),
		H1(site.Title),
		Img(Src("img/office.jpg")),
		P("Notes by ", site.Author),

		H2("Start"),
		Ul(
			site.AddPage("Start", gettingStartedWithProgramming()),
		),

		H2("Design"),
		Ul(
			site.AddPage("Design", purposeOfFuncMain()),
			site.AddPage("Design", nexusPattern()),
			site.AddPage("Design", gracefulServerShutdown()),
			site.AddPage("Design", componentsDiagram()),
			site.AddPage("Design", strictMode()),
			site.AddPage("Design", roleBasedService()),
		),
		H3("Go packages"),
		Ul(
			gregoryv("draw", "software engineering diagrams"),
			gregoryv("web", "html generation"),
		),

		H2("Test"),
		Ul(
			site.AddPage("Test", inlineTestHelpers()),
			site.AddPage("Test", alternateDesign()),
			site.AddPage("Test", setupTeardown()),
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
			site.AddPage("Build", embedVersionAndRevision()),
		),
		H3("Go packages"),
		Ul(
			gregoryv("stamp", "build information code generator"),
			gregoryv("find", "files by name or content"),
		),

		H2("Drills"),
		Ul(
			site.newDrill("-h", "drill/flag_names.go"),
		),
	)
	index := newPage("index.html", findH1(toc), Header(Code(
		versionField(), " - ", Released(),
	)), toc, Footer())
	site.pages = append(site.pages, index)

	return &site
}

type Website struct {
	Title  string
	Author string
	pages  []*Page
	themes []*CSS
	drills []*Page
}

// AddPage creates a new page and returns a link to it
func (me *Website) AddPage(right string, article *Element) *Element {
	title := findH1(article)
	filename := filenameFrom(title) + ".html"

	page := newPage(
		filename,
		stripTags(title)+" - "+me.Title,
		Header(Code(
			right+" - "+A(Href("index.html"), me.Title).String(),
		)),
		article,
		Footer(me.Author),
	)
	me.pages = append(me.pages, page)
	return linkToPage(page)
}

func (me *Website) AddThemes(v ...*CSS) {
	me.themes = append(me.themes, v...)
}

func (me *Website) newDrill(args string, filename string) *Element {
	article := Article(
		loadExample(filename),
		example(args, filename),
	)
	page := NewFile(filepath.Base(toHtmlFile(filename)),
		Html(Lang("en"),
			Head(
				Meta(Charset("utf-8")),
				Meta(
					Name("viewport"),
					Content("width=device-width, initial-scale=1.0"),
				),
				stylesheet("../theme.css"),
				stylesheet("../a4.css"),
				Title(""),
			),
			Body(
				Header(),
				article,
				Footer(me.Author),
			),
		),
	)
	me.drills = append(me.drills, page)
	return linkDrill(filename)
}

// Saves all pages and table of contents
func (me *Website) SaveTo(base string) error {
	for _, page := range me.pages {
		page.SaveTo(base)
	}
	for _, theme := range me.themes {
		theme.SaveTo(base)
	}
	drills := filepath.Join(base, "drill")
	os.MkdirAll(drills, 0722)
	for _, drill := range me.drills {
		if err := drill.SaveTo(drills); err != nil {
			return err
		}
	}
	return nil
}
