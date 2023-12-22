package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/umeshiscreative/reset-service/application"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel()

	fmt.Println("Server is starting...")

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("Unable to serve the Server")
	}
}
