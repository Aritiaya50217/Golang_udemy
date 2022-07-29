package eng

import  "testing"

func TestGreet(t *testing.T){
	want := "Hello"
	if got := Greet(); got != want {
		t.Errorf("Hello() = %q,want : %q",got,want)
	} 
}