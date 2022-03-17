package context

import (
	"context"
	"fmt"
	"time"
)

func doSomethingCool(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timed out")
			err := ctx.Err()
			fmt.Println(err)
			return
		default:
			fmt.Println("doing something cool")
		}
		time.Sleep(500 * time.Millisecond)
	}

}

func Run() {
	fmt.Println("Go Context Tutorial")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go doSomethingCool(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("oh no, I've exceeded the deadline")
	}

	time.Sleep(2 * time.Second)
}

//func enrichContext(ctx context.Context) context.Context {
//	return context.WithValue(ctx, "api-key", "my-secret-key")
//}
//
//func doSomethingCool(ctx context.Context) {
//	apiKey := ctx.Value("api-key")
//	fmt.Println(apiKey)
//	apiname := ctx.Value("api-name")
//	fmt.Println(apiname)
//}
//
//func Run() {
//	fmt.Println("Go Context Tutorial")
//	ctx := context.Background()
//	ctx = enrichContext(ctx)
//	doSomethingCool(ctx)
//}
