package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "Foo", 1)
	ctx = context.WithValue(ctx, "Bar", 2)

	fmt.Println(ctx.Value("Foo").(int)) // 1
	fmt.Println(ctx.Value("Bar").(int)) // 2
}
