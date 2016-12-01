package main

import (
    "fmt"
)

func sort(sortarr []int, value int) {
    sortlen := len(sortarr)
    for i := 0; i < sortlen; i++ {
        if sortarr[i] < value {
            for j := sortlen-1; j > i; j-- {
                sortarr[j] = sortarr[j-1]
            }
            sortarr[i] = value
            break
        }
    }
}

func main() {
    list := []int {2, 5, 6, 3, 9, 7, 19}
    sortarr := make([]int, len(list), len(list))

    for _, num := range(list) {
        sort(sortarr, num)
    }

    fmt.Println(sortarr)
}
