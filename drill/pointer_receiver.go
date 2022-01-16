// Pointer receiver or not
package drill

func init() {
	n := Name{first: "John", last: "Vincic"}
	n.SetFirst("Gregory")
	n.SetLast("Doe") // missing pointer receiver, works on copy

	println(n.String())
}

type Name struct {
	first string
	last  string
}

func (n *Name) SetFirst(v string) { n.first = v }
func (n Name) SetLast(v string)   { n.last = v }

func (me Name) String() string {
	return me.first + " " + me.last
}
