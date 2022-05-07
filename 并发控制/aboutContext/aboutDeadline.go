package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 3000 * time.Millisecond

func main() {
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}

}
