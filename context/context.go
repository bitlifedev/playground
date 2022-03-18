package context

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func doSomethingCool(ctx context.Context) {
	for {
		rand.Seed(time.Now().UnixNano())
		r1 := rand.Intn(20)
		select {
		case <-ctx.Done():
			fmt.Println("timed out")
			err := ctx.Err()
			fmt.Println(err)
			return

		default:
			fmt.Println("Working " + strconv.Itoa(r1))
		}
		if r1 == 10 {
			fmt.Println("Done!")
			return
		}
		time.Sleep(1 * time.Second)
	}

}

func Run() {
	fmt.Println("Go Context Tutorial")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go doSomethingCool(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("oh no, I've exceeded the deadline")
	}
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
