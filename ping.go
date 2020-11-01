package main

import (
        "fmt"
        "net"
        "github.com/sparrc/go-ping"
        "time"
)

func dockerOnline(host string) (bool) {
        timeout := time.Duration(20) * time.Millisecond
        conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, "2375"), timeout)
        if err != nil {
            return false
        }
        if conn != nil {
            defer conn.Close()
       //     fmt.Println("Opened", net.JoinHostPort(host, port))
	    return true
        }
            return false
}

func pingTest(ip string) (bool, time.Duration) {
        pinger, err := ping.NewPinger(ip)

        if err != nil {
                fmt.Printf("ERROR: %s\n", err.Error())
                return false, 0
        } else {
                pinger.SetPrivileged(true)
                pinger.Count = 1
                pinger.Timeout = time.Duration(20) * time.Millisecond

                pinger.Run() // blocks until finished
                stats := pinger.Statistics()

                if pinger.PacketsRecv > 0 {
                        return dockerOnline(ip), stats.AvgRtt
                }
        }
        return false, 0
}

