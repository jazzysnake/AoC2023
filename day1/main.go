package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "errors"
)

func returnFirstAndLastDigits1(input string) (int, error) {
    firstFound := false
    secondFound := false
    var first int
    var second int
    for _, char := range input {
        i, err := strconv.Atoi(string(char))
        if err == nil{
            if !firstFound {
                first = i
                firstFound = true
            } else {
                second = i
                secondFound = true
            }
        }
    }  
    if !firstFound {
        return 0, errors.New("String contained no digits")
    }
    if !secondFound {
        str := fmt.Sprintf("%d%d", first, first)
        res, _ := strconv.Atoi(str)
        return res, nil
    }
    str := fmt.Sprintf("%d%d", first, second)
    res, _ := strconv.Atoi(str)
    return res, nil
}

func returnFirstAndLastDigits2(input string) (int, error) {
    firstFound := false
    secondFound := false
    var first int
    var second int
    for _, char := range input {
        i, err := strconv.Atoi(string(char))
        if err == nil{
            if !firstFound {
                first = i
                firstFound = true
            } else {
                second = i
                secondFound = true
            }
        }
        // TODO implement logic for finding digits as 
        // strings like "one", "two" etc
    }  
    if !firstFound {
        return 0, errors.New("String contained no digits")
    }
    if !secondFound {
        str := fmt.Sprintf("%d%d", first, first)
        res, _ := strconv.Atoi(str)
        return res, nil
    }
    str := fmt.Sprintf("%d%d", first, second)
    res, _ := strconv.Atoi(str)
    return res, nil
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // optionally, resize scanner's capacity for lines over 64K, see next example
    sum:=0
    for scanner.Scan() {
        i, err := returnFirstAndLastDigits1(scanner.Text())
        if err == nil {
            sum += i
        } else {
            log.Print(err)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    println(sum)
}
