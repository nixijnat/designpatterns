package singleton

import (
	"sync"
)

type singleton struct {
	name string
}

func (s *singleton) Do() {
	// TODO
}

var (
	l    *singleton
	once sync.Once
)

/* starve mode

func init()  {
	l = &singleton{"test"}
}

func Get() *singleton {
	return l
}
*/

// lazy mode
func Get() *singleton {
	once.Do(func() {
		l = &singleton{"test"}
	})
	return l
}
