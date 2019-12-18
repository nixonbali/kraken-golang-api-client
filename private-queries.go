package main

import (
  "time"
)

func Nonce() int64 {
  return time.Now().UnixNano()
}
