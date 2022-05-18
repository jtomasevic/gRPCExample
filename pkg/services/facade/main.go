package main

import (
	"fmt"
	"os"

	"github.com/jtomasevic/gRPCExample/pkg/services/facade/app"
)

func main() {
	if err := app.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
