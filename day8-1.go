package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "strings"
    "strconv"
)

type Instruction struct {
    action string
    arg int
}

func (instr Instruction) run(indice int, acc int) (int, int) {
    if instr.action == "acc" {
        indice = indice + 1
        acc = acc + instr.arg
    }
    if instr.action == "jmp" {
        indice = indice + instr.arg
    }
    if instr.action == "nop" {
        indice = indice + 1
    }
    return indice, acc
}

func main() {
    file, err := os.Open("input8.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = file.Close(); err != nil {
            log.Fatal(err)
        }
    }()
    scanner := bufio.NewScanner(file)
    var pile[]Instruction
    for scanner.Scan() {
        input := scanner.Text()
        groupsText := strings.Split(input, " ")
        if len(groupsText) != 2 {
            log.Fatal("Split unsuccessful")
        }

        arg, err := strconv.Atoi(groupsText[1])
        if err != nil {
            log.Fatal(err)
        }
        instruction := Instruction{
            action: groupsText[0],
            arg: arg,
        }
        pile = append(pile, instruction)
    }
    answ := runPile(pile)
    fmt.Println(answ)
}

func runPile(pile []Instruction) int{
    accumulator := 0
    indice := 0
    var collectedIndices []int
    for {
        indice, accumulator = pile[indice].run(indice, accumulator)
        if intInSlice(collectedIndices, indice) {
            break
        } else {
            collectedIndices = append(collectedIndices, indice)
        }
    }

    return accumulator
}

func intInSlice(list []int, a int) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
