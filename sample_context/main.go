package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// Background() と TODO()
	ctxBackground := context.Background()
	ctxTODO := context.TODO()
	fmt.Println("Background context: ", ctxBackground)
	fmt.Println("TODO context: ", ctxTODO)

	// WithCancel()
	ctxCancel, cancel := context.WithCancel(ctxBackground)
	go func() {
		select {
		case <-ctxCancel.Done():
			fmt.Println("WithCancel: context canceled")
		}
	}()
	time.Sleep(500 * time.Millisecond)
	cancel() // キャンセル
	time.Sleep(500 * time.Millisecond)

	// WithTimeout()
	ctxTimeout, cancelTimeOut := context.WithTimeout(ctxBackground, 2*time.Second)
	defer cancelTimeOut()
	go func() {
		select {
		case <-ctxTimeout.Done():
			fmt.Println("WithTimeout: context timeout")
		}
	}()
	time.Sleep(3 * time.Second)

	// WithValue()
	f := func(ctx context.Context) {
		if v := ctx.Value("key"); v != nil {
			fmt.Println("WithValue: key=", v)
		} else {
			fmt.Println("WithValue: key not found")
		}
	}
	ctxValue1 := context.WithValue(ctxBackground, "key", "value")
	go f(ctxValue1)
	ctxValue2 := context.WithValue(ctxBackground, "key2", "value")
	go f(ctxValue2)

	// Wait for goroutine to complete
	time.Sleep(1 * time.Second)
}
