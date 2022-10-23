package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Printf("=========")
	p, _ := filepath.Abs("..")
	fmt.Printf(p)
}
