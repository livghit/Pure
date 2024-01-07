package main

import (
	"context"
	"os"

	"github.com/livghit/TemplIcons/heroicons"
)

func main() {
	icon := heroicons.ServerStack()
	icon.Render(context.Background(), os.Stdout)
}
