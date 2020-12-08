package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "strings"
    "strconv"
    "time"
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
    start := time.Now()

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

    ter := false
    pile_ := make([]Instruction, len(pile))
    var answ int
    i := 0
    for ;!ter;i++ {
        copy(pile_, pile)
        pile_ = updatePile(pile_, i)
        answ, ter = runPile(pile_)
    }
    fmt.Println(answ)

    fmt.Println("Execution duration: " + time.Now().Sub(start).String())
}

func updatePile(pile []Instruction, ind int) []Instruction {
    i := 0
    for j, instruct := range pile {
        if instruct.action == "jmp" || instruct.action == "nop"  {
            i++
            if i == ind {
                if instruct.action == "jmp" {
                    pile[j].action = "nop"
                } else if instruct.action == "nop" {
                    pile[j].action = "jmp"
                }
                return pile
            }
        }
    }
    return pile
}

func runPile(pile []Instruction) (int, bool) {
    accumulator := 0
    indice := 0
    terminated := false
    var collectedIndices []int
    for {
        if indice >= len(pile) {
            terminated = true
            break
        }
        indice, accumulator = pile[indice].run(indice, accumulator)
        if intInSlice(collectedIndices, indice) {
            break
        } else {
            collectedIndices = append(collectedIndices, indice)
        }
    }

    return accumulator, terminated
}

func intInSlice(list []int, a int) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
