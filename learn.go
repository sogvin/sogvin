package sogvin

import (
	. "github.com/gregoryv/web"
)

var gettingStartedWithProgramming = Article(
	H1("Getting started with programming in Go"),

	P(`Programming involves reading and writing code, you do this in
       an text editor. There are plenty to choose from. Below is a
       short list of editors to choose from. The important thing is,
       <em>it has to be a text editor</em>, not a word processor like
       Word.  `),

	Ul(
		Li(A(Href("https://wiki.gnome.org/Apps/Gedit"), "GEdit")),
		Li(A(Href("https://atom.io/"), "Atom")),
		Li(A(Href("https://notepad-plus-plus.org/"), "Notepad++")),
		Li(A(Href("https://www.sublimetext.com/"), "Sublime")),
	),

	P(`Download the latest version from`,
		A(Href("https://golang.org/dl"), "https://golang.org/dl"),
		`. Follow the `, A(Href("https://golang.org/doc/install"),
			`installation instructions`), `carefully and make sure it
		works.`),

	H2("Styling conventions"),

	P(`On these pages; commands are indicated with a gray background
	   with a darker left border.`),

	ShellCommand("go version"),

	`Source code is presented in two ways; an entire files have a gray
      background with a solid border`,

	LoadGoFile("internal/learn/main.go", 0, -1),

	"whereas partial code is without borders.",

	LoadGoFile("internal/learn/main.go", 3, -1),

	H2("The terminal"),

	P(`Once a program is written you'll want to execute it. Without a
       graphical user interface you do this in a terminal. These are
       small programs which enable you to type commands and when you
       hit the enter key they are executed. The go compiler that you
       installed earlier is such a command. If you followed the
       installation instructions thoroughly, you should already have
       used the terminal so I assume you know where to find it.`))
