package main

import (
	"net"
	"../check"
)

func main() {
	port := "16385"
	l, err := net.Listen("tcp", ":" + port)
	check.Check(err)
	defer l.Close()
	println("Listening on port " + port)
	for {
		c, err := l.Accept()
		check.Check(err)
		go handleConnection(c)
	}
}



func handleConnection(conn net.Conn) {
    defer conn.Close()

    // Les data fra tilkoblingen inn i buffer
    buffer := make([]byte, 1024)
		 // rlen vil ha informasjon om hvor langt den har lest
	  rlen, err := conn.Read(buffer)
		check.Check(err)

		// Vi gjør bufferdata om til en string
		msg := string(buffer[0:rlen])
		println(msg)

		// Send returmelding
		//conn.Write([]byte("Affirmative: " + msg))
}
