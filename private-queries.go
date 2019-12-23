package main

import (
  "time"
)

type PostData struct {
  Nonce int64 `json:"nonce"`
}

func Nonce() int64 {
  return time.Now().UnixNano()
}
