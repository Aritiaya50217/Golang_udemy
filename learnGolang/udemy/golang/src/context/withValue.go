package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	ctx2 := context.WithValue(ctx, "a", "1")
	ctx3 := context.WithValue(ctx2, "b", "2")
	ctx4 := context.WithValue(ctx3, "c", "3")

	lookup(ctx, "ctx :", "a")
	lookup(ctx2, "ctx2 :", "a")
	lookup(ctx3, "ctx3 :", "b")
	lookup(ctx3, "ctx3 :", "c")
	lookup(ctx4, "ctx4 :", "c")

	/*
		ctx : <nil> => ไม่มี value
		ctx2 : 1 => "a" มีค่าเท่ากับ "1"
		ctx3 : 2 => "b" มีค่าเท่ากับ "2"
		ctx3 : <nil>  => เป็น nil เพราะ ctx3 รับค่า "b"
		ctx4 : 3 => "c"  มีค่าเท่ากับ "3"
	*/

}

func lookup(ctx context.Context, name, k string) {
	fmt.Println(name, ctx.Value(k))
}
