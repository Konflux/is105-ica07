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
  // Vi åpner en enkel connection til localhost på port 16384
  conn, err := net.Dial("udp", "127.0.0.1:16384")
  check(err)
  defer conn.Close()
  // Så sender vi en hemmelig melding.
  conn.Write([]byte("Møte Fr 5.5 14:45 Flåklypa"))
}
