package main

import (
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Please input addr.  ./server 1.1.1.1:8081")
	}
	localAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:8999") // 固定穿孔内网IP
	if err != nil {
		log.Fatalln("Can't resolve localAddr: ", err)
	}

	remotelyAddr, err := net.ResolveUDPAddr("udp", os.Args[1]) // 远程拨号IP
	if err != nil {
		log.Fatalln("Can't resolve remotelyAddr: ", err)
	}

	udp, err := net.DialUDP("udp", localAddr, remotelyAddr)
	if err != nil {
		log.Fatalln("Udp Dial Err: ", err)
	}

	//

}
