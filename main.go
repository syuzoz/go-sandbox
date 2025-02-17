package main

import (
    "fmt"
    "strconv"
)

func main() {
    for i := 1; i <= 100; i++ {
        fmt.Println(fizzbuzz(i))
    }
}

func fizzbuzz(i int) string {
    if i%3 == 0 && i%5 == 0 {
        return "fizzbuzz"
    } else if i%3 == 0 {
        return "fizz"
    } else if i%5 == 0 {
        return "buzz"
    } else {
        return strconv.Itoa(i)
    }
}
