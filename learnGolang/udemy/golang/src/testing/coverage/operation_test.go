package coverage

import (
	"testing"
)

func TestBasicOperation(t *testing.T) {
	x := operation(1, 2, '+')
	if x != 3 {
		t.Errorf("operation(%f , %f , +) = %f.Wanted: %f", 1.0, 2.0, x, 3.0)
	}
	x := operation(1, 2, '-')
	if x != -1 {
		t.Errorf("operation(%f , %f, - ) = %f.Wanted: %f", 1.0, 2.0, x, -1.0)
	}
	x := operation(1, 2, '*')
	if x != 2 {
		t.Errorf("operation(%f , %f , *) = %f.Wanted: %f", 1.0, 2.0, x, 2.0)
	}
	x := operation(1, 2, '/')
	if x != 0.5 {
		t.Errorf("operation(%f , %f , /) = %f.Wanted: %f", 1.0, 2.0, x, 0.5)
	}
}
