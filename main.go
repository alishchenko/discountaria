package main

import (
	"github.com/alishchenko/discountaria/internal/cli"
	"os"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
