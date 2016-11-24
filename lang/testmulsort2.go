package main

import (
    "fmt"
    "sort"
)

// A couple of type definitions to make the units clear.
type earthMass float64
type au float64

// A Planet defines the properties of a solar system object.
type Planet struct {
    name     string
    mass     earthMass
    distance au
}

type By func(p1 , p2 *Planet) bool

/**
 * 这个是最有意思的地方
 * 比较函数自身绑定排序方法处理排序
 */
func (by By) Sort(planets []Planet) {
    ps := &planetSorter {
        planets: planets,
        by: by,
    }
    sort.Sort(ps)
}

// planetSorter joins a By function and a slice of Planets to be sorted.
type planetSorter struct {
    planets []Planet
    by      func(p1, p2 *Planet) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *planetSorter) Len() int {
    return len(s.planets)
}

// Swap is part of sort.Interface.
func (s *planetSorter) Swap(i, j int) {
    s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *planetSorter) Less(i, j int) bool {
    return s.by(&s.planets[i], &s.planets[j])
}

var planets = []Planet{
    {"Mercury", 0.055, 0.4},
    {"Venus", 0.815, 0.7},
    {"Earth", 1.0, 1.0},
    {"Mars", 0.107, 1.5},
}

// ExampleSortKeys demonstrates a technique for sorting a struct type using programmable sort criteria.
func main() {
    // Closures that order the Planet structure.
    name := func(p1, p2 *Planet) bool {
        return p1.name < p2.name
    }
    mass := func(p1, p2 *Planet) bool {
        return p1.mass < p2.mass
    }
    distance := func(p1, p2 *Planet) bool {
        return p1.distance < p2.distance
    }
    decreasingDistance := func(p1, p2 *Planet) bool {
        return !distance(p1, p2)
    }

    // Sort the planets by the various criteria.
    //转换类型为By
    By(name).Sort(planets)
    fmt.Println("By name:", planets)

    By(mass).Sort(planets)
    fmt.Println("By mass:", planets)

    By(distance).Sort(planets)
    fmt.Println("By distance:", planets)

    By(decreasingDistance).Sort(planets)
    fmt.Println("By decreasing distance:", planets)

}
