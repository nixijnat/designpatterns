package memento

// 直接备份对象
// 注意深拷贝和浅拷贝问题
type mementoObj struct {
	m OriginatorObj
}

type OriginatorObj struct {
	State string
}

func (o *OriginatorObj) createMemento() OriginatorObj {
	return OriginatorObj{o.State}
}

func (o *OriginatorObj) restoreMemento(m OriginatorObj) {
	o.State = m.State
}
