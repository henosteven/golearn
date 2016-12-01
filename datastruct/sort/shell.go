package main

import (
    "fmt"
)

func sort(arr []int) {
   var j, gap int
   var len int = len(arr)
   for gap = len/2; gap > 0; gap /= 2 { //每次循环缩小gap
        for j = gap; j < len; j++ {
            if arr[j] < arr[j - gap] { //相隔gap的两个数字比较, 如果后面的数字比前面的数字小
                tmp := arr[j]
                k := j - gap

                //将当前数字取出来，前面的数字开始往后挪位置，每次跨度为gap
                for k >= 0 && arr[k] > tmp {
                    arr[k + gap] = arr[k]
                    k -= gap
                }

                //该挪位置元素已经好了，把取出的数字放回去
                arr[k + gap] = tmp
            }
        }
   }
}

func main() {
    list := []int {6, 7, 3, 2, 9, 10}
    fmt.Println(list)
    sort(list)
    fmt.Println(list)
}
