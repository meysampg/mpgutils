package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"notify"
	"strings"
	"time"
)

var verbose = flag.Bool("v", false, "Verbose status of connection")
var dialTimeout = flag.Int("t", 10, "Timeout of sytem on dial host in seconds")
var domain = flag.String("d", "google.com", "Critiria doamin for checking connectivity")
var port = flag.String("p", "80", "Port of host which you wanna dial with it")
var cycleTimeout = flag.Int("c", 1, "Cycle time to check connectivity in seconds")

func main() {
	flag.Parse()
	var isConnected bool = false

	if *verbose {
		var info string = fmt.Sprintf("Check connectivity on %s:%s (with %ds for host response timeout) every %ds.", *domain, *port, *dialTimeout, *cycleTimeout)
		log.Print(contifStr(info))
	}

	for {
		conn, err := net.DialTimeout(
			"tcp",
			strings.Trim(*domain, " ")+":"+strings.Trim(*port, " "),
			time.Duration(*dialTimeout)*time.Second)

		if err != nil && isConnected {
			isConnected = false
			notify.Show(":(", "Net is Down!")

			if *verbose {
				log.Print(contifErr("Internet is disconnected", err))
			}
		} else if err == nil {
			conn.Close()

			if !isConnected {
				isConnected = true
				notify.Show(":)", "Net is UP!")

				if *verbose {
					log.Print(contifStr("Internet is connected"))
				}
			}
		}

		if *verbose {
			log.Print(contifStr("Sleep... zZzZz"))
		}
		time.Sleep(time.Duration(*cycleTimeout) * time.Second)
	}
}

func contifStr(str string) string {
	return fmt.Sprintf("[contif] %s", str)
}

func contifErr(msg string, err error) error {
	return fmt.Errorf("%s (Error: %s)", contifStr(msg), err)
}
