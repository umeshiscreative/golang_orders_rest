package main

import (
	"fmt"

	"github.com/umeshiscreative/reset-service/application"
)

func main() {
	app := application.New()

	err := app.Start()
	if err != nil {
		fmt.Println("Unable to serve the Server")
	}
}
