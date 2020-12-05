package main

import "fmt"
import "os"
import "bufio"
import "log"
import "strconv"
import "strings"

func main() {
    file, err := os.Open("input2.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = file.Close(); err != nil {
            log.Fatal(err)
        }
    }()
    var input []string
    scanner := bufio.NewScanner(file)
    valid := 0
    for scanner.Scan() {
        var line = scanner.Text()
        input = strings.Split(line, ":")

        rules := strings.Split(input[0], " ")
        targetLetter := rules[1]
        targetPositions := strings.Split(rules[0], "-")
        targetPos1, _ := strconv.Atoi(targetPositions[0])
        targetPos2, _ := strconv.Atoi(targetPositions[1])
        count := 0
        // No need to decrement 1. There is a space at the begining of each string
        if string(input[1][targetPos1]) == targetLetter  {
            count++
        }
        if string(input[1][targetPos2]) == targetLetter  {
            count++
        }
        if count == 1 {
            valid ++
        }
    }

    fmt.Println(valid)
}
