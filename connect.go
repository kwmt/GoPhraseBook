package main

import (
	"fmt"
	"net"
	"strconv"
)

func tryConnect(network, host string, port int) net.Conn {
	p:= strconv.Itoa(port)
	addr := net.JoinHostPort(host, p)
	fmt.Printf("%v\n",addr)
	c, e := net.Dial(network, addr)
	fmt.Printf("%v\n",c)
	if e == nil { return c }
	return nil
}

func connect(network, service, host string) net.Conn {
	_, addrs, _ := net.LookupSRV(service, network, host)
	fmt.Printf("%v\n",addrs)
	for _, srv := range addrs {
		c := tryConnect(network, srv.Target, int(srv.Port))
		if c != nil {
			return c
		}
	}
	port, _ := net.LookupPort(network, service)
	fmt.Printf("%v\n",port)
	ips, _ := net.LookupHost(host)
	fmt.Printf("%v\n",ips)
	for _, ip := range ips {
		c:= tryConnect(network, ip, port)
		if c != nil {
			return c
		}
	}
	return nil
}

func main() {
	c := connect("tcp", "http", "informit.com")
	fmt.Printf("%v\n",c)
	c.Write([]byte("GET / HTTP/1.1\r\nHost: informit.com\r\n\r\n"))
	buffer := make([]byte, 1024)
	c.Read(buffer)
	fmt.Printf("%s", buffer)
}