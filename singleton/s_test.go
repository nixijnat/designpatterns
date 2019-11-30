package singleton

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	one := Get()
	two := Get()
	fmt.Printf("pointors : %x , %x", one, two)
	if one != two {
		t.Errorf("error: %x, %x", one, two)
	}
}
