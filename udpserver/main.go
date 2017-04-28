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
  port := "16384" // Porten vi ønsker at serveren skal høre etter UDP-pakker
  log.Printf("Starting server, listening on port %s\n", port )
  // Vi hører på alle innkommende forespørsler ved å si at vi hører på
  // :16384.
  addr, err := net.ResolveUDPAddr("udp", ":" + port)
  check(err)
  // Åpne en UDP socket
  sock, err := net.ListenUDP("udp", addr)
  check(err)

  for {
    buf := make([]byte, 1024)
    // ReadFromUDP blokkerer tråden til en melding kommer inn
    // og putter data inn i buf, samt returnerer en read length som vi bruker
    // senere
    rlen, _, err := sock.ReadFromUDP(buf)
    check(err)
    // Vi printer ut buffer fra 0 til read length, og ikke lengre enn det.
    log.Println(string(buf[0:rlen]))
  }
}
