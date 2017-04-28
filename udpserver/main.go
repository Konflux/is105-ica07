package main

import (
  "net"
  "log"
)

func check(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func main() {
  port := "16384"
  log.Printf("Starting server, listening on port %s\n", port )
  addr, err := net.ResolveUDPAddr("udp", ":" + port)
  check(err)
  sock, err := net.ListenUDP("udp", addr)
  check(err)

  for {
    buf := make([]byte, 1024)
    rlen, _, err := sock.ReadFromUDP(buf)
    check(err)
    log.Println(string(buf[0:rlen]))
  }
}
