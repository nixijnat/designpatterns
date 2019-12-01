package memento

import "reflect"

type caretakerTotal struct {
	bak map[string]interface{}
}

type OriginatorTotal struct {
	State1 string
	State2 string
	State3 string
}

func (o *OriginatorTotal) createMemento() map[string]interface{} {
	return backup(o)
}

func (o *OriginatorTotal) restoreMemento(m map[string]interface{}) {
	restore(o, m)
}

/// util
// o is ptr
func backup(o interface{}) map[string]interface{} {
	typ := reflect.TypeOf(o).Elem()
	val := reflect.ValueOf(o).Elem()
	num := typ.NumField()
	m := make(map[string]interface{}, num)
	for i := 0; i < num; i++ {
		m[typ.Field(i).Name] = val.Field(i).Interface()
	}
	return m
}

// o is ptr
func restore(o interface{}, m map[string]interface{}) {
	// typ := reflect.TypeOf(o)
	val := reflect.ValueOf(o).Elem()
	for key, tmp := range m {
		val.FieldByName(key).Set(reflect.ValueOf(tmp))
	}
}
