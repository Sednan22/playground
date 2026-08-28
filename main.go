package main

import (
	"context"
	"fmt"
	"time"
)

type TestResponse struct {
	Value bool
	Err   error
}

func main() {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "val")

	start := time.Now()
	response := someContext(ctx)
	if response.Err != nil {
		fmt.Println(response.Err)
		return
	}
	fmt.Printf("Value come from fetchsomething = %v\n", response.Value)
	fmt.Println("It took: ", time.Since(start))
}

func someContext(ctx context.Context) TestResponse {

	valFromCtx := ctx.Value("key")
	fmt.Println("Value come from ctx =", valFromCtx)

	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	respch := make(chan TestResponse)

	go func() {
		val, err := fetchsomething()
		respch <- TestResponse{
			Value: val,
			Err:   err,
		}
	}()

	select {
	case <-ctx.Done():
		return TestResponse{Err: fmt.Errorf("fetch took to long.")}
	case resp := <-respch:
		return TestResponse{
			Value: resp.Value,
			Err:   resp.Err,
		}
	}

}

func fetchsomething() (bool, error) {
	time.Sleep(10 * time.Millisecond)
	return false, nil
}
