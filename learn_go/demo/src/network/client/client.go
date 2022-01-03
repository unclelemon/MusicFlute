package main

import (
    "io"
    "log"
    "net"
    "os"
)

func main() {
    tcpAddr, err := net.ResolveTCPAddr("tcp6", "[2408:8440:d40:4cd5:80e:58e7:1ae1:88ac]:8000")
    if err != nil {
        log.Fatal(err)
    }
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        log.Fatal(err)
    }
    done := make(chan struct{})
    go func() {
        io.Copy(os.Stdout, conn) 
        log.Println("done")
        done <- struct{}{} 
    }()
    mustCopy(conn, os.Stdin)
    conn.CloseWrite()
    <-done 
}

func mustCopy(dst io.Writer, src io.Reader) {
    if _, err := io.Copy(dst, src); err != nil {
        log.Fatal(err)
    }
}