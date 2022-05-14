package testing

import "testing"

// func Test[Name](*testing.T){}

// func testAdd(t *testing.T) {
// 	a := add(1, 2)
// 	if a != 3 {
// 		t.Error("a := add(1,2) is not 3")
// 	}
// }

func testAddComplex(t *testing.T) {
	tests := []struct {
		lhs      int
		rhs      int
		expected int
	}{
		{lhs: 1, rhs: 2, expected: 3},
	}
	for _, test := range tests {
		ans := add(test.lhs, test.rhs)
		if ans != test.expected {
			t.Errorf("add(%d , %d) = %d .Wanted %d", test.lhs, test.rhs, test.expected)
		}
	}
}
