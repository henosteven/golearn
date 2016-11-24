package main

import (
    "fmt"
    "net"
)

func main() {
    server := "localhost:5000"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
    
    if err != nil {
        fmt.Println("fatal error:", err.Error())
    }

    conn, err := net.DialTCP("tcp", nil, tcpAddr)

    if err != nil {
        fmt.Println("fatal error", err.Error())
    }

    words := "hello world"
    conn.Write([]byte(words))
}
