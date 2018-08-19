package main

import (
    "fmt"
    "strconv"
)

func main() {
    var x string
    var y string

    fmt.Print("What is the first number? ")
    fmt.Scanf("%s", &x)
    actualX, _ := strconv.Atoi(x)

    if actualX < 0 {
        fmt.Println("양수만 입력 받습니다.")
        return
    }

    fmt.Print("What is the second number? ")
    fmt.Scanf("%s", &y)
    actualY, _ := strconv.Atoi(y)
    if actualY < 0 {
        fmt.Println("양수만 입력 받습니다.")
        return
    }

    fmt.Println(
        fmt.Sprintf("%d + %d = %d\n", actualX, actualY, plus(actualX, actualY)) +
            fmt.Sprintf("%d - %d = %d\n", actualX, actualY, minus(actualX, actualY)) +
            fmt.Sprintf("%d * %d = %d\n", actualX, actualY, multi(actualX, actualY)) +
            fmt.Sprintf("%d / %d = %d\n", actualX, actualY, devide(actualX, actualY)))
}

func devide(actualX int, actualY int) int {
    return actualX / actualY
}

func multi(actualX int, actualY int) int {
    return actualX * actualY
}

func minus(actualX int, actualY int) int {
    return actualX - actualY
}

func plus(actualX int, actualY int) int {
    return actualX + actualY
}
