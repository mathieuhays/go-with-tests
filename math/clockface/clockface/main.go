package main

import (
	"os"
	"time"

	"hello/math/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
