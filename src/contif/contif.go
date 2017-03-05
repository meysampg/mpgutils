package main

import (
	"notify"
	"time"
	"net"
	"log"
	"fmt"
)

func main() {
	var isConnected bool = false
	
	for {
		conn, err := net.DialTimeout("tcp", "google.com:80", 1 * time.Second)
		if err != nil && isConnected {
			isConnected = false
			notify.Show(":(", "Net is Down!")
			log.Print(contifStr("Internet is disconnected"))
		} else if err == nil && !isConnected {
			isConnected = true
			notify.Show(":)", "Net is UP!")
			log.Print(contifStr("Internet is connected"))
		}
		
		if err == nil {
			conn.Close()
		}
		
		log.Print(contifStr("Sleep... zZzZz"))
		time.Sleep(5 * time.Second)
	}
}

func contifStr(err string) string {
	return fmt.Sprintf("[contif] %s", err)
}