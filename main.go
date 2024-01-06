package main

import (
	"context"
	"os"

	"github.com/livghit/templkit/components"
)

func main() {
	component := components.Index()
	component.Render(context.Background(), os.Stdout)
}
