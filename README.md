# P2PNetWork
P2PNetWork 区块网络通讯基础 P2P网络实现  (UDP打洞  升级TCP)
``` 

                        Server S
                    207.148.70.129:9981
                           |
                           |
    +----------------------|----------------------+
    |                                             |
  NAT A                                         NAT B
120.27.209.161:6000                            120.26.10.118:3000
    |                                             |
    |                                             |
 Client A                                      Client B
  10.0.0.1:9982                                 192.168.0.1:9982

突破NAT
```
实现原理: 

Server 开启UDP服务

Client 与 Server 建立连接  此时NAT设备 建立 端口映射

NAT映射  之上建立TCP通讯

