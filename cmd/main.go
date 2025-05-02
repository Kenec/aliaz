package main

import (
	"context"
	"log"
	"os"

	"github.com/Kenec/aliaz/cmd/aliaz/root"
)

func main() {
	cmd := root.Commands()

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
