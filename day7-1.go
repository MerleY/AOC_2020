package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "regexp"
)

type InsideBag struct {
    inside map[string]int
}

func main() {
    file, err := os.Open("input7.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = file.Close(); err != nil {
            log.Fatal(err)
        }
    }()

    scanner := bufio.NewScanner(file)
    //bags := make(map[string]InsideBag)
    for scanner.Scan() {
        input := scanner.Text()

        re := regexp.MustCompile(`^(\w+\s){1,2}bags\scontain(\s(\d)\s(\w+\s){1,2}bags(,|\.))+$`)
        matched := re.FindAllSubmatch([]byte(input), -1)
        for i := 0; i < len(matched); i++ {
            for j := 0; j < len(matched[i]); j++ {
                // for k := 0; k < len(matched[i][j]); k++ {
                    fmt.Println(string(matched[i][j]))
                // }
            }
        }
        return
    }
}
