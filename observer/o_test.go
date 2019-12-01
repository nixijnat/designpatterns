package observer

import "testing"

func TestObserver(t *testing.T) {
	// simple test
	var loba = observable{name: "didadi"}
	loba.Add(&observer{"first"})
	var sec = &observer{"second"}
	loba.Add(sec)
	loba.Add(&observer{"third"})
	loba.Notify()
	loba.Remove(sec)
	loba.Notify()
}
