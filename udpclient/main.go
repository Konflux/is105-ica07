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
  conn, err := net.Dial("udp", "127.0.0.1:16384")
  check(err)
  defer conn.Close()
  conn.Write([]byte("Møte Fr 5.5 14:45 Flåklypa"))
}
