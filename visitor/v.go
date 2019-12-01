package visitor

import "log"

type visitor interface {
	visit1(v visitee1)
	visit2(v visitee2)
}

type visitee interface {
	Accept(v visitor)
}

type visitee1 struct{}

func (v visitee1) Accept(vr visitor) {
	vr.visit1(v)
}

type visitee2 struct{}

func (v visitee2) Accept(vr visitor) {
	vr.visit2(v)
}

type visitor1 struct{}

func (vr visitor1) visit1(v visitee1) {
	log.Println("visit1")
}

func (vr visitor1) visit2(v visitee2) {
	log.Println("visit2")
}

func example() {
	var vtor visitor1
	var vtees = []visitee{visitee1{}, visitee2{}}
	for i := range vtees {
		vtees[i].Accept(vtor)
	}
}
