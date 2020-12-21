package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	day21()
	fmt.Println("Execution duration: " + time.Now().Sub(start).String())
}
