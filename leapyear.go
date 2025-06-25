package main

import (
    "fmt"
)

func leapyear(year int) bool {
    return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func main() {
    var year int
    fmt.Print("请输入一个年份：")
    fmt.Scan(&year)

    if leapyear(year) {
        fmt.Printf("%d 是闰年。\n", year)
    } else {
        fmt.Printf("%d 不是闰年。\n", year)
    }
}