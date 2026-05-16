package main

import "fmt"

func add(a, b int) int {
    return a + b
}

func main() {
    num1, num2 := 1, 4
    result := add(num1, num2)
    fmt.Printf("func(%d, %d) = %d\n", num1, num2, result)
}