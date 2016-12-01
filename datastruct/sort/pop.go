package main

import (
    "fmt"
)

/*
 * 每次判断相邻两个值，交换顺序
 * 每次冒出一个最大值
 */
func sort(list []int) {
    len := len(list)
    for j := 0; j < len; j++ {
        for i := 0; i < len - 1; i++ {
            if list[i] > list[i+1] {
                list[i] , list[i+1] = list[i+1] , list[i]
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
    list := []int {6, 7, 3, 2, 9}
    display(list)
    fmt.Println("====")
    sort(list)
    display(list)
}
