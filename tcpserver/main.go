package main

import (
	"net"
	"log"
)

func main() {
	l, err := net.Listen("tcp", ":16385")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		c, err := l.Accept()
		if err!= nil {
			log.Fatal(err)
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
    //some code...
    
    //Simple read from connection
    buffer := make([]byte, 1024)
	  rlen, err := c.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(buffer[0:rlen]))
}