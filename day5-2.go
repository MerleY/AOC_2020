package main

import "fmt"
import "os"
import "bufio"
import "log"
import "math"

func main() {
    file, err := os.Open("input5.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = file.Close(); err != nil {
            log.Fatal(err)
        }
    }()

    scanner := bufio.NewScanner(file)
    var passportsId []int
    for scanner.Scan() {
        input := scanner.Text()
        var id int
        if len(input) != 10 {
            log.Fatal("line with wrong length", len(input))
        }

        row := 0
        col := 0
        for i, letter := range input {
            if i < 7 {
                if letter == 'B' {
                    row = row + int(math.Pow(2, float64(6 - i)))
                }
            } else {
                if letter == 'R' {
                    col = col + int(math.Pow(2, float64(9 - i)))
                }
            }
        }
        id = row * 8 + col
        passportsId = append(passportsId, id)
    }

    for i := 7; i < 894; i++ {
        if !intInSlice(passportsId, i) {
            fmt.Println(i)
        }
    }
}

func intInSlice(list []int, a int) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
