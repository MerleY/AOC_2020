package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	day17()
	fmt.Println("Execution duration: " + time.Now().Sub(start).String())
}
