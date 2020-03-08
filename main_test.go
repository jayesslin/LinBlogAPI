package main

import (
	"os"
	"testing"
)

func TestIntegerStuff(t *testing.T) {
	print(os.Getenv("GOOS"), "\n")
}
