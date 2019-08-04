package main

import (
    "fmt"
    "net"
    "io"
    "os"
)

func main(){
    listener, err := net.Listen("tcp",":8080")

    if err != nil {
        // error
    }

    defer listener.Close()
    for {
        conn, err := listener.Accept()

        if err != nil {
            // handle error
        }

        fmt.Println("Starting transfer")

        go func(c net.Conn) {
            io.Copy(c, os.Stdout)
            fmt.Println("Connection closing")
            c.Close()
        }(conn)

        fmt.Println("%T", conn)
    }
}
