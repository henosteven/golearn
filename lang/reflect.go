package main

import (
    "fmt"
    "reflect"
)

type MyInt int

type User struct {
    name string
    age int
}

func (user User) ShowName() int {
    fmt.Println("---", user.name, "----")
    return 0
}

func (user User) SetName(name string) {
    user.name = name
}

func (user User) JustCall(str string) {
    fmt.Println("hello ", str)
}

func main() {

    var user = User{"heno", 27}
    
    tp := reflect.TypeOf(user)
    vf := reflect.ValueOf(user)

    fmt.Println(tp) //type *main.User
    fmt.Println(tp.NumMethod())
    fmt.Println(tp.Method(0)) // {SetName  func(*main.User, string) <func(*main.User, string) Value> 0}
    fmt.Println(tp.Method(0).Name, tp.Method(0).Type) //SetName func(main.User, string)

    fmt.Println(tp.Name()) // User
    fmt.Println(tp.Kind()) // Struct
    fmt.Println(tp.NumField())
    fmt.Println(tp.Field(0))//{name main string  0 [0] false}
    fmt.Println(tp.Field(0).Name, tp.Field(0).Type) // name string

    vf.MethodByName("ShowName").Call([]reflect.Value{}) //--- heno ----
    vf.MethodByName("JustCall").Call([]reflect.Value{reflect.ValueOf("henosteven")}) //hello henosteven

    var i int = 10
    fmt.Println("type:", reflect.TypeOf(i)) // type: int
    fmt.Println("value:", reflect.ValueOf(i)) // value: 10

    r := reflect.ValueOf(i)
    fmt.Println(r.Type()) // int
    fmt.Println(r.Int()) // 10
    fmt.Println(r.Kind()) // int

    var j MyInt  = 11
    r2 := reflect.ValueOf(j)
    fmt.Println(r2.Type()) // main.MyInt
    fmt.Println(r2.Kind()) // int

    var x float64 = 3.4
    fmt.Println(reflect.ValueOf(x))
}
