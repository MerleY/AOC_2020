package main

import "fmt"
import "os"
import "bufio"
import "log"
import "strconv"

func main() {
    file, err := os.Open("input1-1.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = file.Close(); err != nil {
            log.Fatal(err)
        }
    }()
    var input []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        i, _ := strconv.Atoi(scanner.Text())
        input = append(input, i)
    }

    for i, n := range input {
        for j, n2 := range input[i:] {
            for _, n3 := range input[j:] {
                if n + n2 + n3 == 2020 {
                    fmt.Println(n2 * n *n3)
                }
            }
        }
    }
}
