package main

import "fmt"
import "os"
import "bufio"
import "log"

func main() {
    file, err := os.Open("input3.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = file.Close(); err != nil {
            log.Fatal(err)
        }
    }()
    scanner := bufio.NewScanner(file)
    tree := 0
    x := 0
    for scanner.Scan() {
        var line = scanner.Text()
        lengthString := len(line)
        relativPos := x % lengthString
        if (string)(line[relativPos]) == "#" {
            tree++
        }
        x = 3 + x
    }

    fmt.Println(tree)
}
