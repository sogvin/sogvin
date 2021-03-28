package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/nexus"
	"github.com/gregoryv/wolf"
)

func main() {
	run(wolf.NewOSCmd())
}

func run(cmd wolf.Command) {
	var (
		cli    = cmdline.New(cmd.Args()...)
		domain = cli.Option("-d, --domain").String("booking")
		name   = cli.Option("-n, --name").String("bookme")
	)

	log.SetFlags(0)

	_, err := os.Stat(domain)
	if err == nil {
		log.Fatalf("directory %q exists", domain)
	}

	os.MkdirAll(filepath.Join(domain, "cmd", name), 0755)

	// main.go
	var buf bytes.Buffer
	p, _ := nexus.NewPrinter(&buf)
	p.Println("package main")
	p.Println("func main() {}")
	writeGoFile(
		filepath.Join(domain, "cmd", name, "main.go"),
		buf.Bytes(),
	)

	// service.go
	buf.Reset()
	p.Println("package ", domain)
	p.Println("type Service struct{}")
	p.Println("func (me *Service) Use(role Role) {")
	p.Println("var u user")
	p.Println("u.setService(me)")
	p.Println("role.setUser(&u)")
	p.Println("}")
	writeGoFile(
		filepath.Join(domain, "service.go"),
		buf.Bytes(),
	)

	// role.go
	buf.Reset()
	p.Println("package ", domain)
	p.Println("type Role struct {")
	p.Println("*user")
	p.Println("}")
	writeGoFile(
		filepath.Join(domain, "role.go"),
		buf.Bytes(),
	)

	// user.go
	buf.Reset()
	p.Println("package ", domain)
	p.Println("type user struct {")
	p.Println("srv *Service")
	p.Println("}")
	p.Println("func (me *user) setService(v *Service) { me.srv = v }")
	writeGoFile(
		filepath.Join(domain, "user.go"),
		buf.Bytes(),
	)

	// resource.go
	buf.Reset()
	p.Println("package ", domain)
	p.Println("// Define resource types for this domain")
	writeGoFile(
		filepath.Join(domain, "resource.go"),
		buf.Bytes(),
	)

}

func writeGoFile(filename string, content []byte) {
	tidy, err := format.Source(content)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(filename, tidy, 0644)

	log.Println(strings.Repeat("-", 40), filename)
	log.Println(string(tidy))
}
