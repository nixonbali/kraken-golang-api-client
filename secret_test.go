package main

import (
  "testing"
  "time"
)

func TestNonce(t *testing.T) {
  nonce0 := Nonce()
  time.Sleep(time.Nanosecond)
  nonce1 := Nonce()
  if nonce1 <= nonce0 {
    t.Error("Nonce not increasing by the nanosecond")
  }
}
