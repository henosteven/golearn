package main

import (
    "fmt"
    "net"
)

func main() {
    netListen, err := net.Listen("tcp", "localhost:5000")
    defer netListen.Close()

    if err != nil {
        fmt.Println("listen failed\n")
    }

    for {
        conn, err := netListen.Accept()
        if err != nil {
            continue
        }

        handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    buffer := make([]byte, 2048)
    for {
        n, err := conn.Read(buffer)
        if err != nil {
            return
        }

        fmt.Println("recive data string:\n", buffer[:n])
    }
}
