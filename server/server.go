package main

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"p2pnetwork/define"
	"sync"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Please input addr.  ./server 0.0.0.0:8081")
	}
	addr, err := net.ResolveUDPAddr("udp", os.Args[1])
	if err != nil {
		log.Fatalln("Can't resolve address: ", err)
	}
	udp, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalln("Udp ListenUDP Error: ",err)
	}

	// 服务ID，服务地址     当建立UDP通讯时 NAT会开启映射(映射稳定情况 网络提供商息息相关)
	var addrMap map[string]*net.UDPAddr
	var addrMu sync.Mutex
	buf := make([]byte,1024)
	for {
		i, udpAddr, err := udp.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		data := buf[:i]
		discovery := &define.Discovery{}
		err = json.Unmarshal(data, discovery)
		if err != nil {
			log.Println(err)
			continue
		}

		addrMu.Lock()
		dis,b := addrMap[discovery.ID]
		addrMu.Unlock()
		if !b {
			addrMap[discovery.ID] = udpAddr
			continue
		}
		if discovery.Request == "" {
			continue
		}
		addrMu.Lock()
		req,b := addrMap[discovery.Request]
		addrMu.Unlock()
		if !b {
			log.Printf("Addr: %s  no ex\n",discovery.Request)
			continue
		}

		disAddr,err := json.Marshal(&define.ServerAddr{Addr: dis.String()})
		if err != nil {
			log.Println("DisAddr Marshal Err")
			continue
		}
		reqAddr,err := json.Marshal(&define.ServerAddr{Addr: req.String()})
		if err != nil {
			log.Println("ReqAddr Marshal Err")
			continue
		}
		_, err = udp.WriteToUDP(reqAddr, dis)
		if err != nil {
			log.Println(err)
			continue
		}

		_, err = udp.WriteToUDP(disAddr,req)
		if err != nil {
			log.Println(err)
			continue
		}

	}
}
