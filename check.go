package main

import (
	"fmt"
	"net"
	"time"

)

func check(destination string, port string) string {
    // Create a new TCP connection
    address := destination + ":" + port
    timeout := time.Duration(5 * time.Second)
    conn, err := net.DialTimeout("tcp", address, timeout)
    var status string
    if err != nil {
        status = fmt.Sprintf("[DOWN] %v is unreachable, \nError: %v", destination, err)
    } else {
        status = fmt.Sprintf("[UP] %v is reachable\nFrom: %v\nTo: %v", destination, conn.LocalAddr(), conn.RemoteAddr())
        conn.Close()
    }
    return status
}