package observer

import "log"

type Observable interface {
	Add(o Obserber)
	Remove(o Obserber)
	Notify()

	Name() string
}

type Obserber interface {
	Do(o Observable, args ...interface{}) error
}

type observer struct {
	name string
}

func (o *observer) Do(ob Observable, args ...interface{}) error {
	log.Printf("%s recieved from %s", o.name, ob.Name())
	return nil
}
func (o *observer) Name() string {
	return o.name
}

type observable struct {
	name  string
	obers []Obserber
}

func (o *observable) Add(ob Obserber) {
	o.obers = append(o.obers, ob)
}

func (o *observable) Remove(ob Obserber) {
	for i := range o.obers {
		if o.obers[i] == ob { // NOTE , == not a good idea
			o.obers = append(o.obers[0:i], o.obers[i+1:]...)
			return
		}
	}
}

func (o *observable) Notify() {
	for i := range o.obers {
		o.obers[i].Do(o, "test")
	}
}

func (o *observable) Name() string {
	return o.name
}
