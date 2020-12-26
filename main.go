package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	day25()
	fmt.Println("Execution duration: " + time.Now().Sub(start).String())
}
