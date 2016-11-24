package main

import (
    "fmt"
)

func modify (s []string) {
    s[1] = "henosteven";
}

func modify2(s [2]string) {
    s[1] = "henosteven";
}
func main() {

    /**
     * 通过下面的输出就更能说明一个 slice是指向underlying数组的指针 
     * 改变slice会改变数组
     */
    
    s := []string {"xiaopang", "jinjing"}
    fmt.Println(s) //[xiaopang jinjing]
    modify(s)
    fmt.Println(s) //[xiaopang henosteven]

    s1 := [2]string {"xiaopang", "jinjing"}
    fmt.Println(s1) //[xiaopang jinjing]
    modify2(s1)
    fmt.Println(s1) //[xiaopang jinjing]
}
