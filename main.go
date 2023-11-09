package main

import (
	"fmt"
	"os"

	"github.com/anfernee/sanitizer/pkg/app"
)

func main() {
	if err := app.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
