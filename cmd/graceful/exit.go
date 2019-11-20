package main

func (me *MainEntry) Exit() int {
	if me.err != nil {
		return 1
	}
	return 0
}
