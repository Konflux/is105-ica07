package main

import (
  "net"
  "../check"
)



func main() {
  port := "16384" // Porten vi ønsker at serveren skal høre etter UDP-pakker
  // Vi hører på alle innkommende forespørsler ved å si at vi hører på
  // :16384.
  addr, err := net.ResolveUDPAddr("udp", ":" + port)
  check.Check(err)
  // Åpne en UDP socket
  sock, err := net.ListenUDP("udp", addr)
  check.Check(err)

  println("Listening on port " + port)

  for {
    buf := make([]byte, 2048)
    // ReadFromUDP blokkerer tråden til en melding kommer inn
    // og putter data inn i buf, samt returnerer en read length som vi bruker
    // senere
    rlen, addr, err := sock.ReadFromUDP(buf)
    check.Check(err)
    // Vi lagrer bufferet fra 0 til read length, som string.
    msg := string(buf[0:rlen])
    println(msg)

    // Vi sender tilbake en godkjenningsmelding
    sock.WriteToUDP([]byte("Affirmative: " + msg), addr)
  }
}
