package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	day18()
	fmt.Println("Execution duration: " + time.Now().Sub(start).String())
}
