package tools

import (
	"log"
	"net"
)

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println("get addr failed|err=", err)
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				log.Println("local ip addr=", ipnet.IP.String())
				return ipnet.IP.String()
			}
		}
	}
	log.Println("failed to get local ip|addrs=", addrs)
	return ""
}

func BlockMain() {
	listen, err := net.Listen("tcp", "localhost:20000")
	if err != nil {
		log.Println("listen failed, err:", err)
		return
	}
	//如果监听成功，则无限循环，监听是否有链接
	for {
		_, err = listen.Accept() //接受listen所监听的端口传入的链接
		if err != nil {
			log.Println("accept failed , err:", err)
			continue
		}
		//没有错误，链接建立完成，处理连接的请求。
	}
}
