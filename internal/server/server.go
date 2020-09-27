package server

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/arrwhidev/go-redis/internal/request"
)

func Start() {
	l, err := net.Listen("tcp4", ":6379")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer l.Close()

	ip := getLocalIP()
	fmt.Println("go-redis is listening at", ip+":6379")

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal("Failed to accept", err)
			return
		}

		r := request.NewRequest(c)
		go r.Handle()
	}
}

func getLocalIP() string {
	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if strings.HasPrefix(ip.String(), "192.168") {
				return ip.String()
			}
		}
	}

	return ""
}
