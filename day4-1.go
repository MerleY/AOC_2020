package main

import "fmt"
import "os"
import "bufio"
import "log"
import "strings"

func main() {
    file, err := os.Open("input4.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = file.Close(); err != nil {
            log.Fatal(err)
        }
    }()
    scanner := bufio.NewScanner(file)
    var passports []string
    line := ""
    for scanner.Scan() {
        input := scanner.Text()
        if input == "" {
            passports = append(passports, line)
            line = ""
        } else {
            line = line + " " + input
        }
    }
    var mandatoryFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
    valid := 0
    for _, passport := range passports {
        var elements = strings.Split(passport, " ")
        var presentFields []string
        for _, element := range elements {
            var keyValue = strings.Split(element, ":")
            key := keyValue[0]
            if stringInSlice(mandatoryFields, key) && !stringInSlice(presentFields, key){
                presentFields = append(presentFields, key)
            }
        }
        if (len(presentFields) > 6) {
            valid++
        }



    }
    fmt.Println(valid)
}
func stringInSlice(list []string, a string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
