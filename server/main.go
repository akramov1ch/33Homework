package main

import (
    "io"
    "net"
    "os"
)

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }
    defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        return
    }
    filename := string(buffer[:n])
    file, err := os.Create(filename)
    if err != nil {
        return
    }
    defer file.Close()
    io.Copy(file, conn)
}
