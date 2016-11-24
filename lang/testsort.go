package main

import (
    "fmt"
    "sort"
)

type Person struct {
    Name string
    Age int
}

func (p Person) String() string {
    return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person
func (a ByAge) Len() int { return len(a)}
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i]}
func (a ByAge) Less(i, j int) bool {return a[i].Age < a[j].Age}

func main() {
    //尽管这么写没有问题，而且下面不用进行类型转换ByAge()了，但是很明显的一个可阅读性就下降了
    //代码好不好，就这就是区别
    //people := ByAge{ 
    people := []Person{
            {"Bob", 31},
            {"John", 42},
            {"Michael", 17},
            {"Jenny", 26},
    }

    fmt.Println(people)
    sort.Sort(ByAge(people))
    fmt.Println(people)
}