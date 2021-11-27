package main

import (
	"os"
	"time"

	"maths" // REPLACE THIS!
)

func main() {
	t := time.Now()
	maths.SvgWriter(os.Stdout, t)
}
