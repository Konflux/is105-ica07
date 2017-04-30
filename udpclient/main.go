package main

import (
  "net"
  "../check"
)


func main() {
  // Vi åpner en enkel connection til localhost på port 16384
  conn, err := net.Dial("udp", "127.0.0.1:16384")
  check.Check(err)
  defer conn.Close()
  // Så sender vi en hemmelig melding.
  conn.Write([]byte("Møte Fr 5.5 14:45 Flåklypa"))

  // Vi vil også lese meldingen vi får i retur
  buf := make([]byte, 1024)
  // conn.Read() vil i dette tilfellet alltid lese over EOF, så
  // vi hopper over errorhandling her.
  n, _ := conn.Read(buf)

  println(string(buf[:n]))
}
