package sogvin

import (
	"os"
	"path/filepath"

	. "github.com/gregoryv/web"
)

func NewWebsite() *Website {
	title := "Software Engineering - Skills &amp; Practice"
	author := "Gregory Vin&ccaron;i&cacute;"
	site := Website{
		title:  title,
		author: author,
	}
	site.AddThemes(a4(), theme())

	toc := Article(Class("toc"),
		H1(title),
		Img(Src("img/office.jpg")),
		P("Notes by ", author),

		H2("Preface"),

		Blockquote("Practice makes perfect"),

		P(`In software engineering there are many skills required to
        become professional. Here I present some of those skills and
        ways to practice them.`),

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

		H2("Test"),
		Ul(
			site.AddPage("Test", inlineTestHelpers()),
			site.AddPage("Test", alternateDesign()),
			site.AddPage("Test", setupTeardown()),
		),

		H2("Build"),
		Ul(
			site.AddPage("Build", embedVersionAndRevision()),
		),

		H2("Drills"),

		P(`Drills are short examples for practicing often used
		concepts.`),

		H3("Command line"),
		Ul(
			site.AddDrill("Command line", "", "drill/flag_types.go"),
			site.AddDrill("Command line", "", "drill/flag_names.go"),
			site.AddDrill("Command line", "", "drill/cmdline_basic.go"),
		),

		H3("Reading files"),
		Ul(
			site.AddDrill("Reading files", "", "drill/openfile.go"),
			site.AddDrill("Reading files", "", "drill/slurp_file.go"),
			site.AddDrill("Reading files", "", "drill/readfile_byline.go"),
		),

		H3("Logging"),
		Ul(
			site.AddDrill("Logging", "", "drill/logging.go"),
			site.AddDrill("Logging", "", "drill/level_logs.go"),
		),

		H3("Encoding"),
		Ul(
			site.AddDrill("Encoding", "", "drill/json_encode.go"),
		),

		H2("References"),
		Ul(
			site.AddPage("References", packageRefs()),
		),
	)
	index := newPage("index.html", findH1(toc), Header(Code(
		versionField(), " - ", Released(),
	)), toc, Footer())
	site.pages = append(site.pages, index)

	return &site
}

type Website struct {
	title  string
	author string
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
		stripTags(title)+" - "+me.title,
		Header(Code(
			right+" - "+A(Href("index.html"), me.title).String(),
		)),
		article,
		Footer(me.author),
	)
	me.pages = append(me.pages, page)
	return linkToPage(page)
}

func (me *Website) AddDrill(right, args string, filename string) *Element {
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
				// todo link to parent
				Header(
					Header(Code(
						right+" - "+A(Href("../index.html"), me.title).String(),
					)),
				),
				article,
				Footer(me.author),
			),
		),
	)
	me.drills = append(me.drills, page)
	return linkDrill(filename)
}

func (me *Website) AddThemes(v ...*CSS) {
	me.themes = append(me.themes, v...)
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
