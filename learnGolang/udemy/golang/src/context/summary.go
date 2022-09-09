package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	ctx2 := context.WithValue(ctx, "b", "1")
	ctx3 := context.WithValue(ctx2, "b", "2")
	ctx4 := context.WithValue(ctx3, "c", "3")

	// http.Request
	lookup(ctx, "ctx :", "a")
	lookup(ctx2, "ctx2 :", "b")
	lookup(ctx3, "ctx3 :", "a")
	lookup(ctx3, "ctx3 :", "b")
	lookup(ctx4, "ctx4 :", "b")

}

func lookup(ctx context.Context, name, k string) {
	fmt.Println(name, ctx.Value(k))
}
