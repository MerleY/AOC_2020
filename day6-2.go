package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
)

type GroupAnswers struct {
    L []byte
    N int
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
    ga.N++
}

func (ga *GroupAnswers) Merge(letters []byte) {
    if ga.N == 0 {
        ga.AddUniqs(letters)
        return
    }

    var temp []byte
    for _, element := range ga.L {
        if byteInSlice(letters, element) {
            temp = append(temp, element)
        }
    }
    ga.L = temp
    ga.N++
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
            fmt.Println(group.Count())
            total = total + group.Count()
            group = GroupAnswers{}
            continue;
        }
        group.Merge([]byte(input))
    }
    fmt.Println(total)

}

func byteInSlice(list []byte, a byte) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
