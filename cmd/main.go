package main

import (
	"context"

	"github.com/vix993/shakemon/internal/api"
)

func main() {
	ctx := context.Background()

	// Start API
	api.Run(ctx)
}
