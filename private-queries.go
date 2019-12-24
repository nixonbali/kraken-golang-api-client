package main

import (
  "time"
  "crypto/hmac"
  "crypto/sha512"
  "crypto/sha256"
  "strconv"
  "net/http"
  "encoding/base64"
  "fmt"
  "io/ioutil"
)

var client = &http.Client{}

type PostData struct {
  Nonce int64 `json:"nonce"`
}

func Nonce() int64 {
  return time.Now().UnixNano()
}

func postAccountBalance() {
  url := "https://api.kraken.com/0/private/Balance"
  uri := "/0/private/Balance"
  req, err := http.NewRequest("POST", url, nil)
  if err != nil {
    panic(err)
  }
  req.Header.Add("API-Key", APIKEY)
  signature := signAPI(uri, "")
  req.Header.Add("API-Sign", signature)
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }
  fmt.Println(resp.Body)

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }
  fmt.Println(string(body))
}

func signAPI(uri string, postData string) string {
  nonce := Nonce()
  postData = "nonce=" + strconv.FormatInt(nonce, 10) + "&" + postData

  // Calculate the SHA256 of the nonce and the POST data
  sha := sha256.New()
  sha.Write([]byte(postData))
  message := sha.Sum(nil)

  // Decode the API secret (the private part of the API key) from base64
  decoded, err := base64.StdEncoding.DecodeString(PRIVATEKEY)
  if err != nil {
    panic(err)
  }

  // Calculate the HMAC of the URI path and the SHA256, using SHA512 as the
  // HMAC hash and the decoded API secret as the HMAC key
  mac := hmac.New(sha512.New, decoded)
  mac.Write(message)
  sum := mac.Sum(nil)


  // Encode the HMAC into base64
  return base64.StdEncoding.EncodeToString(sum)
}
