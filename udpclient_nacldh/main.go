package main

import (
  "golang.org/x/crypto/nacl/box"
  "crypto/rand"
  "../check"
  "net"
)

func main() {

  // Generer ny pubkey og privkey
  pubkey, privkey, err := box.GenerateKey(rand.Reader)
  check.Check(err)

  // sharedKey kommer til å være lik mellom server og klient
  var sharedKey [32]byte

  // Public key fra server
  var serverKey [32]byte

  // Melding som skal krypteres
  msg := []byte("Møte Fr 5.5 14:45 Flåklypa")

  // Nonce for krypteringssalt
  var nonce [24]byte
  noncebuf := make([]byte, 24)

  rand.Reader.Read(noncebuf)

  copy(nonce[:], noncebuf[:24])



  // Vi åpner en enkel connection til localhost på port 16384
  conn, err := net.Dial("udp", "127.0.0.1:16384")
  check.Check(err)
  defer conn.Close()

  // Send publickey til server
  conn.Write(pubkey[:])

  // Motta publickey fra server
  buf := make([]byte, 2048)
  conn.Read(buf)

  // Kopier buffer til clientKey
  copy(serverKey[:], buf[:32])

  // Precompute sharedKey
  box.Precompute(&sharedKey, &serverKey, privkey)

  // Send nonce til server
  conn.Write(nonce[:])
  var out []byte;

  // Generer kryptert melding
  encrypted := box.SealAfterPrecomputation(out, msg, &nonce, &sharedKey)

  // Send kryptert melding til server
  conn.Write(encrypted)

}
