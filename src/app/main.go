package main

import (
	"config"
	"notify"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

const (
	down        = ":("
	downMessage = "Net is Down!"
	up          = ":)"
	upMessage   = "Net is UP!"
)

func main() {
	flag.Parse()
	var isConnected bool = false
	config.Initialize()
	config.SetConfigParameter()
	if *config.Verbose {
		var info string = fmt.Sprintf("Check connectivity on %s:%s (with %ds for host response timeout) every %ds.", *config.Domain, *config.Port, *config.DialTimeout, *config.CycleTimeout)
		log.Print(contifStr(info))
	}

	for {
		conn, err := net.DialTimeout(
			"tcp",
			strings.Trim(*config.Domain, " ")+":"+strings.Trim(*config.Port, " "),
			time.Duration(*config.DialTimeout)*time.Second)

		notif := notify.New(notify.Options{
			DefaultIcon: "icon/default.png",
			AppName:     "Net checker",
		})
		if err != nil && isConnected {
			isConnected = false
			notif.Push(
				down,
				downMessage,
				"",
				notify.UR_NORMAL,
			)

			if *config.Verbose {
				log.Print(contifErr("Internet is disconnected", err))
			}
		} else if err == nil {
			conn.Close()

			if !isConnected {
				isConnected = true
				notif.Push(
					up,
					upMessage,
					"",
					notify.UR_NORMAL)

				if *config.Verbose {
					log.Print(contifStr("Internet is connected"))
				}
			}
		}

		if *config.Verbose {
			log.Print(contifStr("Sleep... zZzZz"))
		}
		time.Sleep(time.Duration(*config.CycleTimeout) * time.Second)
	}
}

func contifStr(str string) string {
	return fmt.Sprintf("[contif] %s", str)
}

func contifErr(msg string, err error) error {
	return fmt.Errorf("%s (Error: %s)", contifStr(msg), err)
}
