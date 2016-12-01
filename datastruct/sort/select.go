package main

import (
    "fmt"
)

/*
 * 每次选出一个最大值
 */
func sort(list []int) {
   len := len(list)
   for i := 0; i < len; i++ {
       for j := 0; j < len; j++ {
           if (list[i] < list[j]) {
                list[i], list[j] = list[j], list[i]
           } 
       }
       fmt.Println(list)
   }
}

func display(list []int) {
    for _, num := range(list) {
        fmt.Println(num)
    }
}

func main() {
    list := []int {6, 7, 3, 2, 9, 13}
    display(list)
    fmt.Println("====")
    sort(list)
    display(list)
}
