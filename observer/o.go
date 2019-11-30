package observer

import "log"

type Observable interface {
	Add(o Obserber)
	Remove(o Obserber)
	Notify()
}

type Obserber interface {
	Update()
}

type observer struct {
	name string
}

func (o *observer) Update() {
	log.Println(o.name, "recieved")
}

type observable struct {
	obers []Obserber
}

func (o *observable) Add(ob Obserber) {
	o.obers = append(o.obers, ob)
}

func (o *observable) Remove(ob Obserber) {
	for i := range o.obers {
		if o.obers[i] == ob {
			o.obers = append(o.obers[0:i], o.obers[i+1:]...)
			return
		}
	}
}

func (o *observable) Notify() {
	for i := range o.obers {
		o.obers[i].Update()
	}
}
