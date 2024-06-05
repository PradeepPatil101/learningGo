package main

import (
	"os"
	"time"

	clockface "learning/mathsClockface/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
