package main

import (
	"context"
	"fmt"

	"github.com/abhirambsn/micro-svc/app"
)

func main() {
	appl := app.New()
	err := appl.Start(context.TODO(), ":3000")
	if err != nil {
		fmt.Println("Failed to start app:", err)
	}
}