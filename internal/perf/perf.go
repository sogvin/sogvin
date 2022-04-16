package main

func main() {
	_ = NewItem()
}

func NewItem() *Item {
	var i Item
	i.x = 2
	return &i
}

type Item struct {
	x int
}
