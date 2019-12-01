package facade

import "log"

// 标准门面模式需要有一个门面类，实际开发过程中
// 可以不用创建一个门面类，直接使用导出的包方法即可
type Facade struct {
	i inserter
	q queryer
}

func (f *Facade) Query() {
	f.q.Do()
}

func (f *Facade) Insert() {
	f.i.Do()
}

type queryer struct{}

func (q *queryer) Do() {
	log.Println("query")
}

type inserter struct{}

func (i *inserter) Do() {
	log.Println("insert")
}
