package main

func (me *MainEntry) skipf(format string, args ...interface{}) bool {
	me.Logf(format, args...)
	if me.err != nil || me.dryrun {
		return true
	}
	return false
}
