package main

import (
    "fmt"
    "net"
    //"sync"
    "os"
    "os/signal"
    "io"
    "log"
    "syscall"
)

func echo_srv(c net.Conn) { //, wg sync.WaitGroup) {
    //wg.Add(1)
    defer c.Close()
    //defer wg.Done()

    msg := make([]byte, 1024)

    for {
        n, err := c.Read(msg)

        if err != nil && err != io.EOF {
            fmt.Printf("ERROR: read\n")
            fmt.Print(err)
            return
        }

        if n != 0 {
            fmt.Printf(" Received: %+v\n", string(msg[:n]))
        }

        if err == io.EOF {
            return
        }

            //fmt.Printf("SERVER: received EOF (%d bytes ignored)\n", n)
        //fmt.Printf("SERVER: received %v bytes\n", n)

        /*n, err = c.Write(msg[:n])
        if err != nil {
            fmt.Printf("ERROR: write\n")
            fmt.Print(err)
            return
        }
        fmt.Printf("SERVER: sent %v bytes\n", n)*/
    }
}

func main() {
    //var wg sync.WaitGroup

    ln, err := net.Listen("unix", "./sock_srv")
    if err != nil {
            fmt.Print(err)
            return
    }
    defer ln.Close()

    sigc := make(chan os.Signal, 1)
    signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGTERM)
    go func(c chan os.Signal) {
        // Wait for a SIGINT or SIGKILL:
        sig := <-c
        log.Printf("Caught signal %s: shutting down.", sig)
        // Stop listening (and unlink the socket if unix type):
        ln.Close()
        os.Remove("./sock_srv")
        // And we're done:
        os.Exit(0)
    }(sigc)

    for {
        conn, err := ln.Accept()
        if err != nil {
                fmt.Print(err)
                continue
        }
        go echo_srv(conn)
    }

}
