package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := time.After(600 * time.Millisecond)

	time.AfterFunc(400*time.Millisecond , func(){
		log.Println("We are about to cancel the ctx")
		cancel()
	})


	select {
	case <-ctx.Done():
		log.Println("ctx.Done()")
		log.Printf("ctx.Err() Done : %v\n", ctx.Err())
	case <-ch:
		log.Println("Ch done")
		log.Printf("ctx.Err() Ch : %v\n", ctx.Err())
	}
}
