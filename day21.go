package main

import (
    "fmt"
    "strings"
    "aoc/input"
    "aoc/arrays"
    "regexp"
    "log"
)

type list struct {
    ingredients []string
}

func (l *list) fusion (l2 []string) {
    if len(l.ingredients) == 0 {
        l.ingredients = l2
    } else {
        tempList := []string{}
        for _, i := range l2 {
            if arrays.StringIn(i, l.ingredients) {
                tempList = append(tempList, i)
            }
        }
        l.ingredients = tempList
    }
}

func day21() {
    input_ := input.Load("21").ToStringArray()
    piba := make(map[string]*list) //piba : potential igredient by allergen
    allIngs := []string{}
    re := regexp.MustCompile(`(.+)\(contains (.+)\)`)
    for _, li := range input_ {
        matches := re.FindAllStringSubmatch(li, -1)
        if len(matches[0]) != 3 {
            log.Fatal("erro")
        }
        ings := strings.Split(matches[0][1], " ")
        for _, ing := range ings {
            if arrays.StringIn(ing, allIngs) {
                allIngs = append(allIngs, ing)
            }
        }
        all := strings.Split(matches[0][2], " ")
        for _, a := range all {
            if _, ok := piba[a]; ok {
                piba[a].fusion(ings)
            } else {
                l := &list{ingredients: ings}
                piba[a] = l
            }
        }
    }
    for _, ing := range allIngs {
        in := false
        for _, list:= range piba {
            if arrays.StringIn(ing, list.ingredients) {
                in = true
                break
            }
        }
        if !in {
            fmt.Println(ing)
        }
    }
}

