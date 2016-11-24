package main

import (
    "fmt"
    "sort"
)

type earthmass float64
type au float64
type Planet struct {
    name string
    mass earthmass
    distance au
}

type planetSorter struct {
    planets [4]Planet
    by func(p1, p2 Planet) bool
}

/**
 * 这里又有一处需要注意的地方
 * 如果这里不传入指针，那肯定是改不了顺序的，因为planets [4]Planet是数组
 * 根据拷贝传值的模式，所以根本不会对原有的数据产生任何影响
 * =======================================================
 * 如果真要改变传值，但是又不想传指针，可以把planets [4]Planet -》 改为 planets []Planet
 * 这样尽管传入的参数不是指针，是拷贝值，但是因为拷贝了指针的值，调整指针指向的值，会修改原值，所以又实现了 
 *
 * 可以试试
 */
func (s *planetSorter) Len() int {
    return len(s.planets)
}

func (s *planetSorter) Swap(i, j int) {
    s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}

func (s *planetSorter) Less(i, j int) bool {
    return s.by(s.planets[i], s.planets[j])
}

var planets = [4]Planet{
    {"Mercury", 550, 4},
    {"Venus", 815, 70},
    {"Earth", 1000, 10},
    {"Mars", 107, 15},
}
    
func main() {
    distance := func(p1, p2 Planet) bool {
        return p1.distance < p2.distance
    }

    s := &planetSorter{planets, distance}
    sort.Sort(s)
    fmt.Println(s)
}
