package main

import (
    "fmt"
    "strconv"
)

const (
    BIN = 2
    DEC = 10
    HEX = 16
)

func convert(baseFrom, baseTo int, input string) string {
    val, _ := strconv.ParseInt(input, baseFrom, 64)
    return strconv.FormatInt(val, baseTo)
}

func main() {
    input := "42"
    fmt.Println(convert(BIN, DEC, "101010"))      // из двоичной в десятичную
    fmt.Println(convert(HEX, BIN, "2a"))           // из шестнадцатеричной в двоичную
    fmt.Println(convert(DEC, HEX, input))          // из десятичной в шестнадцатеричную
}
