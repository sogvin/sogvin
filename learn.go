package sogvin

import (
	. "github.com/gregoryv/web"
)

var gettingStartedWithProgramming = Article(
	H1("Getting started with programming in Go"),
	Boxnote("Pick a text editor", 0.4),
	P(

		`Programming involves writing code, you do this with an text
         editor. There are plenty to choose from. I recommend you
         choose one of the editors listed below depending on your
         platform. Though it's like picking out your first bicycle, it
         should be your decision. Please do a internet search for
         text-editors and try out others aswell. The important thing
         is, <em>it has to be a text editor</em>, not a word processor
         like Word.`,
	),
	Ul(
		Li(
			A(Href("https://wiki.gnome.org/Apps/Gedit"), "GEdit"),
			" - my recommendation if you've never used one before",
		),
		Li(A(Href("https://atom.io/"), "Atom")),
		Li(A(Href("https://notepad-plus-plus.org/"), "Notepad++")),
		Li(A(Href("https://www.sublimetext.com/"), "Sublime")),
	),

	Boxnote("Download Go", 0.1),
	P(

		`Right, the editor up and running. Time to install the Go
		language and its standard packages. Get the latest version
		from`, A(Href("https://golang.org/dl"),
			"https://golang.org/dl"), `. Follow the `,
		A(Href("https://golang.org/doc/install"), `installation
		instructions`), `carefully and make sure it works. In this
		process you will need to use a terminal. The terminal is where
		you will enter commands for the computer to execute. Install
		the language now, it shouldn't take more than a few minutes.`,
	),
)
