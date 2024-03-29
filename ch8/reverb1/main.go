package main

import (
   "log"
   "net"
   "time"
   "strings"
   "fmt"
   "bufio"
)

func main(){
    listener, err := net.Listen("tcp","localhost:8000")
    if err != nil {
        log.Fatal(err)
    }
    for {
        conn ,err := listener.Accept()
        if err != nil {
            log.Print(err)
            continue
        }
        handleConn(conn)
    }
}

func echo(c net.Conn,shout string,delay time.Duration){
    fmt.Fprintln(c,"\t",strings.ToUpper(shout))
    time.Sleep(delay)
    fmt.Fprintln(c,"\t",shout)
    time.Sleep(delay)
    fmt.Fprintln(c,"\t",strings.ToLower(shout))
}

func handleConn(c net.Conn){
    defer c.Close()
    input:= bufio.NewScanner(c)
    for input.Scan() {
        go echo(c, input.Text(), 1*time.Second)
    }
}