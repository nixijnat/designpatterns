package memento

import (
	"log"
	"testing"
)

func TestTotal(t *testing.T) {
	var origin = OriginatorTotal{"1", "2", "3"}
	var taker caretakerTotal
	taker.bak = origin.createMemento()
	log.Println(origin)
	origin.State1, origin.State2, origin.State3 = "4", "5", "6"
	log.Println(origin)
	origin.restoreMemento(taker.bak)
	log.Println(origin)
}
