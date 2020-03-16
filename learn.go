package sogvin

import (
	. "github.com/gregoryv/web"
)

var gettingStartedWithProgramming = Article(
	H1("Getting started with programming in Go"),

	H2("Text editor"),
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
	H2("Install language"),
	P(`Download the latest version from`,
		A(Href("https://golang.org/dl"), "https://golang.org/dl"),
		`. Follow the `, A(Href("https://golang.org/doc/install"),
			`installation instructions`), `carefully and make sure it
		works.`),

	H2("Styling conventions"),

	P(`On these pages; commands are indicated with a gray background
	   with a darker left border. Commands start with the '$'
	   character and other lines are the output if any.`),

	ShellCommand("$ go version\ngo version go1.14 linux/amd64"),

	`Source code is presented in two ways; an entire file has a gray
     background with a solid border.`,

	LoadGoFile("internal/learn/main.go", 0, -1),

	"whereas partial content is without borders.",

	LoadGoFile("internal/learn/main.go", 3, -1),

	H2("Setup working directory"),

	P(`Once a program is written you'll want to execute it. Without a
       graphical user interface you do this in a terminal. The
       terminal is a small program which enables you to type
       commands. When you hit the enter key the command is
       executed. The go compiler that you installed earlier is such a
       command. If you followed the installation instructions
       thoroughly, you should already have used the terminal so I
       assume you know where to find it.`),

	P(`First; use the terminal and create a folder where you will be
       working.`),

	ShellCommand(`$ cd $HOME
$ mkdir go-learn
$ cd go-learn`),

	H2("Your first program"),
	P(`Now create a file with the following content and save it as
	<code>hello_world.go</code>`),

	LoadGoFile("internal/learn/hello_world.xgo", 0, -1),

	P(`What does this program do? execute it with the following command`),
	ShellCommand("$ go run hello_world.go\nHello, world!"),
)
