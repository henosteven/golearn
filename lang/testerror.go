package main

import (
        "fmt"
       )

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    if e < 0 {
        return fmt.Sprintf("cannot Spart negative number")
    } else {
        return ""
    }
}

func Sqrt(x float64) (float64, error) {
    if (x > 0) {
        return 0, nil
    } else {
       return -1, ErrNegativeSqrt(-2)
    }
}

func main() {
    result, err := Sqrt(-2)
    fmt.Println(result, err)
}

