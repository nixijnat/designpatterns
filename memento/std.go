package memento

type memento struct {
	state string
}

type caretaker struct {
	m memento
}

type Originator struct {
	State string
}

func (o *Originator) createMemento() memento {
	return memento{o.State}
}

func (o *Originator) restoreMemento(m memento) {
	o.State = m.state
}

func example() {
	var origin = Originator{"test"}
	var taker caretaker
	taker.m = origin.createMemento()
	origin.restoreMemento(taker.m)
}
