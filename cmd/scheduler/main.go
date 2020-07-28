package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/kinderyj/k8s-scheduler-framework-demo/pkg/plugins"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	cmd := plugins.Register()
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
