package main

import (
	"os"

	"github.com/AgileEngineTest/api"
)

func main() {

	a := api.New()
	err := a.Run()

	if err != nil {
		os.Exit(1)
	}
}
