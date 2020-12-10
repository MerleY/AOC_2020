package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "strconv"
    "time"
    "sort"
)


func main() {
    start := time.Now()
    lines := getFileContent("inputs/input10.txt")
    input := transform(lines)
    sort.Ints(input)

    jolt := 0
    three_joltDelta := 1
    one_joltDelta := 0
    for i := 0; i < len(input); i++ {
        if input[i] == jolt + 1 {
            jolt++
            one_joltDelta++
        }
        if input[i] == jolt + 3 {
            three_joltDelta++
        }
        jolt = input[i] 
    }
    fmt.Println(one_joltDelta * three_joltDelta)
    fmt.Println("Execution duration: " + time.Now().Sub(start).String())
}

func findConsecutivSum(table []int, ind int) []int{
    for start := 0; start < ind; start++ {
        sum := table[start]
        for i := start + 1; i < ind; i++ {
            sum += table[i]
            if sum > table[ind] {
                break
            }
            if sum == table[ind] {
                return table[start:i]
            }
        }
    }
    return nil
}

func sumOfTwoBefore(table []int, sum int) bool {
    if (len(table) != 25) {
        log.Fatal("Table length invalid")
    }
    for i, a := range table {
        for _, b := range table[i:] {
            if a + b == sum {
                return true
            }
        }
    }

    return false
}

func intInSlice(list []int, a int) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func min(list []int) int{
    var min int
    for i, b := range list {
        if i == 0 {
            min = b
        } else if (b < min) {
            min = b
        }
    }
    return min
}

func max(list []int) int{
    var max int
    for i, b := range list {
        if i == 0 {
            max = b
        } else if (b > max) {
            max = b
        }
    }
    return max
}

func transform (lines []string) []int {
    var numbers []int
    for i, line := range lines {
        n, err := strconv.Atoi(line)
        if err != nil {
            log.Fatal("Could not convert line " + string(i))
        }
        numbers = append(numbers, n)
    }

    return numbers
}

func getFileContent(name string) []string {
    var lines []string
    file, err := os.Open(name)
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
        input := scanner.Text()
        lines = append(lines, input)
    }

    return lines
}
