package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	day19()
	fmt.Println("Execution duration: " + time.Now().Sub(start).String())
}
