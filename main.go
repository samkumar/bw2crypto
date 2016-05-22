package main

import (
  "fmt"
  "os"
)

func main() {
  if len(os.Args) != 3 {
    fmt.Printf("Usage: %s <signing key> <verifying key>\n", os.Args[0])
    os.Exit(0)
  }
  sk, e := UnFmtKey(os.Args[1])
  if e != nil {
    fmt.Printf("Could not unformat signing key: %v\n", e)
    os.Exit(1)
  }
  vk, e := UnFmtKey(os.Args[2])
  if e != nil {
    fmt.Printf("Could not unformat verifying key: %v\n", e)
    os.Exit(1)
  }
  if !CheckKeypair(sk, vk) {
    fmt.Println("valid keypair failed to validate")
  }
}
