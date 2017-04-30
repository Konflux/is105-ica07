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
    // Lag ny sharedKey for hver klient
    var sharedKey [32]byte
    // Public key fra klient
    var clientKey [32]byte

    // Nonce fra klient
    var nonce [24]byte

    // Melding
    var message []byte

    // Usikker hva denne brukes til
    var out []byte

    // Buffer for lagring av data fra pakker
    buf := make([]byte, 2048)

    // Motta publickey fra klient
    _, addr, err := sock.ReadFromUDP(buf)
    check.Check(err)

    // Kopier buffer til clientKey
    copy(clientKey[:], buf[:32])


    // Precompute sharedKey
    box.Precompute(&sharedKey, &clientKey, privkey)

    // Send publickey til klient
    sock.WriteToUDP(pubkey[:], addr)

    // Motta nonce fra klient
    _, addr, err = sock.ReadFromUDP(buf)
    check.Check(err)
    copy(nonce[:], buf[:24])

    // Motta hemmelig melding
    rlen, addr, err := sock.ReadFromUDP(buf)
    check.Check(err)
    message = buf[:rlen]

    // Dekrypter meldingen
    decrypted, _ := box.OpenAfterPrecomputation(out, message, &nonce, &sharedKey)
    println(string(decrypted))

  }
}
