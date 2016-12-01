package main

import (
    "fmt"
    "math"
)

type Tree_elt struct {
    parent *Tree_elt
    value []int
    pList []*Tree_elt
    count int
    isleaf bool
}

var MAX = 5 - 1
var MIN = math.Ceil(5 / 2)

func insertNodeList(node *Tree_elt, value int) {
    for i := 0; i < len(node.value); i++ {
        if node.value[i] > value {
            node.value[i], node.value[i+1] = node.value[i+1], node.value[i] 
        }
    }
}

func getInsertPos(node *Tree_elt, value int) int {
    var i int = 0
    for i = 0; i < len(node.value); i++ {
        if node.value[i] > value {
            break
        }
    }
    return i
}

func insertToParent(node *Tree_elt, value int) {

}

func insert(node *Tree_elt, value int) {
    if node == nil {
        node = &Tree_elt{nil, []int{value}, nil, 1, true}    
    } else {
        if node.isleaf {
            if (node.count == MAX) {
                //当前节点分裂
                newNode1 := &Tree_elt{node.parent, []int{value}, nil, 1, true}
                newNode2 := &Tree_elt{node.parent, []int{value}, nil, 1, true}
                i := MAX / 2
                node.parent.pList[i] = newNode1
                node.parent.pList[i+1] = newNode2

                //中间元素向上升级
                if node.parent == nil {
                    node.parent := &Tree_elt{nil, []int{value}, nil, 1, true}
                } else {
                    insertToParent(node.parent, value) 
                }
            } else {
                //插入当前节点 - 移动当前点即可 
                insertNodeList(node, value)
            }
        } else {
            //寻路
            i := getInsertPos(node, value)
            node = node.pList[i]
        }
    }
}

func walk(node *Tree_elt) {
    if node != nil {
        for i := 0; i < node.count; i++ {
            fmt.Println(node.value[i])
            if node.pList[i] != nil {
                walk(node.pList[i])
            }
        }
    }
}

func search(node *Tree_elt, value int) {
   if node == nil {
       fmt.Println("not found")
   } else {
       for i := 0; i < node.count; i++ {
           if node.value[i] == value {
               fmt.Println("found")
               break
           }

           if value < node.value[i] {
               search(node.pList[i+1], value)
           }
       }
   } 
}

func main() {
    fmt.Println("B-tree")
    numList := []int{17, 8, 12, 35, 26, 30, 65, 87, 9, 10, 3, 5, 13, 15, 28, 29, 36, 60, 75, 79, 90, 99}

    var root *Tree_elt
    for _, num := range(numList) {
        insert(root, num)
    }

    walk(root)
}
