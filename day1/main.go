package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func returnFirstAndLastDigits1(input string) (int, error) {
	firstFound := false
	secondFound := false
	var first int
	var second int
	for _, char := range input {
		i, err := strconv.Atoi(string(char))
		if err == nil {
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
	stringDigits := map[string]int{
		"one":   0,
		"two":   0,
		"three": 0,
		"four":  0,
		"five":  0,
		"six":   0,
		"seven": 0,
		"eight": 0,
		"nine":  0,
	}
  stringDigitsList := []string{
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
  }

	for _, char := range input {
		i, err := strconv.Atoi(string(char))
		if err == nil {
			if !firstFound {
				first = i
				firstFound = true
			} else {
				second = i
				secondFound = true
			}
		}
    //println("------", string(char))
		for key, index := range stringDigits {
      stringDigit := []rune(key)
      // chars match
      if stringDigit[index] == char{
        stringDigits[key] = index + 1
        //println("match", key, index + 1)
        // all chars found in a digit
        if index+1 == len(key) {
          if !firstFound {
            first = indexOf(key, stringDigitsList) + 1
            firstFound = true
          } else {
            second = indexOf(key, stringDigitsList) + 1 
            secondFound = true
          }
          stringDigits[key] = 0
          //resetMap(stringDigits)
          //break
        }
      } else if rune(key[0]) == char {
          stringDigits[key] = 1
        } else {
        stringDigits[key] = 0
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

func indexOf(str string, strList []string) int {
  for i, item := range strList {
    if item == str {
      return i
    }
  }
  return -1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
    txt := scanner.Text()
		i, err := returnFirstAndLastDigits2(txt)
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
