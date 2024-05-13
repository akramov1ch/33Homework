package main

import (
    "flag"
    "fmt"
    "io"
    "net"
    "os"
)

func main() {
    flag.Parse()
    if flag.NArg() == 0 {
        fmt.Println("Iltimos, kamida bitta fayl yo'lini kiriting")
        return
    }

    for _, filepath := range flag.Args() {
        sendFile(filepath)
    }
}

func sendFile(filepath string) {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Serverga ulanishda xato:", err)
        return
    }
    defer conn.Close()
    fileInfo, err := os.Stat(filepath)
    if err != nil {
        fmt.Println("Fayl haqida ma'lumot olishda xato:", err)
        return
    }
    _, err = conn.Write([]byte(fileInfo.Name()))
    if err != nil {
        fmt.Println("Fayl nomini yuborishda xato:", err)
        return
    }

    // Faylni ochish
    file, err := os.Open(filepath)
    if err != nil {
        fmt.Println("Faylni ochishda xato:", err)
        return
    }
    defer file.Close()
    _, err = io.Copy(conn, file)
    if err != nil {
        fmt.Println("Fayl ma'lumotlarini yuborishda xato:", err)
    }
}
