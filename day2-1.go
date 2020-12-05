package main

import "fmt"
import "os"
import "bufio"
import "log"
import "strconv"
import "strings"

func main() {
    file, err := os.Open("input2-1.txt")
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
        c := strings.Count(input[1], targetLetter)
        minAndMax := strings.Split(rules[0], "-")
        fmt.Println(minAndMax)
        min, _ := strconv.Atoi(minAndMax[0])
        max, _ := strconv.Atoi(minAndMax[1])
        if c >= min && c <= max {
            valid++
        }
    }

    fmt.Println(valid)
}
