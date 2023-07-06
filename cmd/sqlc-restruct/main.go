package main

import (
	"context"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	if err := App.RunContext(ctx, os.Args); err != nil {
		log.Fatal(err)
	}
}
