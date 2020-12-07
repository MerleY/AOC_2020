package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
)

type GroupAnswers struct {
    L []byte
}

func (ga *GroupAnswers) AddUniq(letter byte) {
    for _, element := range ga.L {
        if element == letter {
            return
        }
    }
    ga.L = append(ga.L, letter)
}

func (ga *GroupAnswers) AddUniqs(letters []byte) {
    for _, letter := range letters {
        ga.AddUniq(letter)
    }
}

func (ga GroupAnswers) Count() int {
    return len(ga.L)
}

func main() {
    file, err := os.Open("input6.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = file.Close(); err != nil {
            log.Fatal(err)
        }
    }()
    scanner := bufio.NewScanner(file)
    total := 0
    group := GroupAnswers{}
    for scanner.Scan() {
        input := scanner.Text()
        if input == "" {
            total = total + group.Count()
            group = GroupAnswers{}
            continue;
        }
        group.AddUniqs([]byte(input))
    }
    fmt.Println(total)

}
