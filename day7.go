package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "strings"
    "regexp"
    "strconv"
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
    bags := make(map[string]InsideBag)
    for scanner.Scan() {
        input := scanner.Text()
        parts := strings.Split(input, " bags contain ")
        bagColor := parts[0]
        bags[bagColor] = InsideBag{inside: make(map[string]int)}
        for _, insideBags := range strings.Split(parts[1], ",") {
            re := regexp.MustCompile(`(\d|no)\s((?:\w+\s)+)bags?\.?`)
            matched := re.FindAllSubmatch([]byte(insideBags), -1)
            color := strings.TrimSpace(string(matched[0][2]))
            if string(matched[0][1]) == "no" {
                continue
            }
            nbrBags, _ := strconv.Atoi(string(matched[0][1]))
            bags[bagColor].inside[color] = nbrBags
        }
    }
    fmt.Println(countContainOut(bags, "shiny gold"))
    fmt.Println(countContainIn(bags, "shiny gold"))
}
func countContainIn(bags map[string]InsideBag, chosenColor string) int {
    total := 0
    for color, nb := range bags[chosenColor].inside {
        total = total + nb * (countContainIn(bags, color) + 1)
    }
    return total
}

func countContainOut(bags map[string]InsideBag, chosenColor string) int {
    var collectBags []string
    collectBags = append(collectBags, chosenColor)
    for {
        var newBags []string
        for color, insideBags := range bags {
            for _, collectBag := range collectBags {
                if insideBags.inside[collectBag] != 0 && !stringInSlice(collectBags, color) && !stringInSlice(newBags, color){
                    newBags = append(newBags, color)
                }
            }
        }
        if len(newBags) == 0 {
            break
        } else {
            collectBags = append(collectBags, newBags...)
        }
    }

    return len(collectBags) - 1
}

func stringInSlice(list []string, a string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
