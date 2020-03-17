package sogvin

import (
	. "github.com/gregoryv/web"
)

var gettingStartedWithProgramming = Article(
	H1("Getting started with programming"),

	P(`Are you an aspiring software developer in need of guidance?
	   then this material is for you. First you'll learn the basics of
	   programming. The goal is for you to understand some of the
	   concepts used while programming. There are plenty of examples
	   for you to learn from along the way. But let's start of by
	   setting up your computer with the minimum required tools for
	   programming.`),

	H2("Editor, compiler and terminal"),

	P(`EDITOR: The first thing you need is a text editor. There are plenty to
       choose from. Below are two editors I think are a good starting
       point. There are lots of others with varying complexity but
       don't start with those unless you're already familiar with
       them. The important thing is, it has to be a text editor, Not a
       word processor like Word.`),

	Ul(
		Li(A(Href("https://wiki.gnome.org/Apps/Gedit"), "GEdit")),
		Li(A(Href("https://notepad-plus-plus.org/"), "Notepad++")),
	),

	P(`COMPILER: You'll be using the Go programming language. It and
		the compiler are available from`,
		A(Href("https://golang.org/dl"), "https://golang.org/dl"),
		`. Follow the `, A(Href("https://golang.org/doc/install"),
			`installation instructions`), `carefully and make sure it
		works.`),

	P(`TERMINAL: Once a program is written you'll want to execute
       it. Without a graphical user interface you do this in a
       terminal. The terminal is a small program which enables you to
       type in commands. When you hit the enter key, the command is
       executed. The go compiler that you installed earlier is such a
       command. If you followed the installation instructions
       thoroughly, you should already have used the terminal so I
       assume you know where to find it.`),

	stylingConventions,
	setupWorkingDirectory,
)

var stylingConventions = Section(
	H2("Styling conventions"),

	P(`On these pages; commands are indicated with a gray background
	   with a darker left border. Commands start after the <code>$</code>
	   character, and other lines are the output of the command. `),

	ShellCommand("$ go version\ngo version go1.14 linux/amd64"),

	P(`Note!  when asked to enter a command do not enter the first
	   <code>$</code> character, only what is after it. Source code is
	   presented in two ways; an entire file has a gray background
	   with a solid border.`),

	LoadFile("internal/learn/main.go"),
	"whereas partial content is without borders.",
	LoadFile("internal/learn/main.go", 3, -1),

	Sidenote("Side note; emphasizing an important concept.", 0.0),
	P(`There is a lot to learn and you whenever a section includes a
	many concepts or longer explanations I'll add a side note with the
	Most important thing. Also the material is formated in such a way
	that if you choose to print it out there is room for your own
	notes on the left hand side.<br> Ok, you have the tools up and
	running and you understand how to read this material. It's time
	to learn programming.`),
)

var setupWorkingDirectory = Section(
	H2("Setup working directory"),

	P(`First; use the terminal and create a folder where you will be
       working.`),

	ShellCommand(`$ cd $HOME
$ mkdir go-learn
$ cd go-learn`),

	H2("Your first program"),
	P(`Now create a file with the following content and save it as
	<code>hello_world.go</code>`),

	LoadFile("internal/learn/hello_world.xgo"),

	P(`What does this program do? execute it with the following command`),
	ShellCommand("$ go run hello_world.go\nHello, world!"),
)
