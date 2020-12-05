package main

import "fmt"
import "os"
import "bufio"
import "log"

func main() {
    factor := 1
    var slopes = [5][2]int {
        {1, 1},
        {3, 1},
        {5, 1},
        {7, 1},
        {1, 2},
    }
    for _ ,slope := range slopes {
        tree := 0
        x := 0
        i := 0
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
        for scanner.Scan() {
            if i % slope[1] != 0 {
                i++
                continue;
            }
            i++
            var line = scanner.Text()
            lengthString := len(line)
            relativPos := x % lengthString
            if (string)(line[relativPos]) == "#" {
                tree++
            }
            x = slope[0] + x
        }
        factor = factor * tree
    }

    fmt.Println(factor)
}
