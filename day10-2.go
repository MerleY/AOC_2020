package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "strconv"
    "time"
    "sort"
)

func main() {
    start := time.Now()
    lines := getFileContent("input10.txt")
    input := transform(lines)
    sort.Ints(input)

    possArrang := make([]int, len(input))
    for i := 0; i <= 3; i++ {
        if input[i] <= 3{
            possArrang[i] = 1
        }
    }

    for i := 0; i < len(input); i++ {
        for j := 1; j <= 3; j++ {
            if i + j < len(input) {
                if input[i + j] <= input[i] + 3 {
                    possArrang[i + j] += possArrang[i]
                }
            }
        }
    }
    fmt.Println(possArrang[len(possArrang) - 1])
    fmt.Println("Execution duration: " + time.Now().Sub(start).String())
}
