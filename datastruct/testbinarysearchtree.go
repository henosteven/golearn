package main

import (
    "fmt"
)

type Tree_elt struct {
    v int
    left *Tree_elt
    right *Tree_elt
}

type Tree struct {
    root *Tree_elt
    count int
}

func (tree *Tree)create() {
    tree.count = 0
    tree.root = nil
}

func (tree *Tree)insert(value int) {
    curNode := tree.root
    tmpNode := &Tree_elt{value, nil, nil}
    if curNode == nil {
        tree.root = tmpNode
    } else {
        for {
           if curNode == nil {
               curNode = tmpNode
               break
           }

           if  value  < curNode.v {
                if curNode.left == nil {
                     curNode.left = tmpNode
                     break
                } else {
                    curNode = curNode.left
                }
           } else {
               if curNode.right == nil {
                    curNode.right = tmpNode
                    break
               } else {
                    curNode = curNode.right 
               }
           }
        }
    }
}

func (tree *Tree)search(value int) bool {
    curNode := tree.root

    for {
        if curNode == nil {
            return false
        } else {
            if curNode.v == value {
                return true
            } else if curNode.v > value {
                curNode = curNode.left
            } else {
                curNode = curNode.right
            }
        }
    }
}

func (tree *Tree)walk(node *Tree_elt) {
    if node != nil {
        tree.walk(node.right)
        fmt.Println(node.v)
        tree.walk(node.left)
    }
}

func (tree *Tree) reverse(node *Tree_elt){
    if node != nil {
        temp := node.left
        node.left = node.right
        node.right = temp

        tree.reverse(node.left)
        tree.reverse(node.right)
    }
}

func main() {
    var tree Tree
    fmt.Println(tree)
    tree.create();
    fmt.Println(tree)

    numList := []int{6, 3, 2, 4, 5, 8, 9, 1}
    for _, num := range(numList) {
        tree.insert(num)
    }

    tree.walk(tree.root)

    if tree.search(9) {
        fmt.Println("found")
    } else {
        fmt.Println("not found")
    }

    tree.reverse(tree.root)
    tree.walk(tree.root)
}
