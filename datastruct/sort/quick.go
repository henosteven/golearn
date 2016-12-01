package main

import (
    "fmt"
)

func sort(s []int, l int, r int) {
    if l < r { //首先判断参数的右边位置>左边位置 
        var i , j, x int = l,  r, s[l]  //取出左边位置的元素
        for i < j  {  
            for i < j && s[j] >= x { // 从右向左找第一个小于x的数  
                j-- 
            }
            if i < j { //如果找到当前左边位置的左边去了就应该放弃-》说明右边没有找到
                s[i] = s[j]; //找到了小于x的数，把该数字放到i槽位 
                i++ //i槽位已有序，故向后移动一位
            }

            for i < j && s[i] < x { // 从左向右找第一个大于等于x的数  
                i++    
            }
            if i < j { //如果找到当前右边位置的右边去了，说明左边没有找到 
                s[j] = s[i]; //找到了大于x的数，把该数字放到之前空出来的j槽位 
                j-- //j槽位有序，故向前移动一位
            }
        }  
        s[i] = x; //i和j已经碰撞在一起 
        sort(s, l, i - 1); 
        sort(s, i + 1, r);  
    } 
}

func display(list []int) {
    for _, num := range(list) {
        fmt.Println(num)
    }
}

func main() {
    list := []int {6, 7, 3, 2, 9, 10}
    display(list)
    fmt.Println("====")
    sort(list, 0, 4)
    display(list)
}
