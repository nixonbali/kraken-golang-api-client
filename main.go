// kraken-api exploration

package main

import (
  "fmt"
  "encoding/json"
  "net/http"
  //"time"
  //"bytes"
  "strings"
)

type serverTime struct {
  Error []string
  Result *serverTimeResult
}

type serverTimeResult struct {
  Unixtime int
  Rfc1123 string
}


type assetInfo struct {
  Error []string
  Result *map[string]asset
}

type asset struct {
  Aclass string
  Altname string
  Decimals int
  DisplayDecimals int `json:"display_decimals"`
}

type assetPairInfo struct {
  Error []string
  Result *map[string]asset
}

// interface for kraken api: some basic struct w/ Error, Result, and APICall methods?
// reference smaller classes?
// defined constants for assets and asset pairs
// i.e. XXBT = "XXBT" to make lookups easier
// cache map for lookups + sotring only asset info we need w/ quick function call to add more

// type assetPair struct {
//     Altname string // alternate pair name
//     wsname string// WebSocket pair name (if available)
//     AclassBase string // asset class of base component
//     base // asset id of base component
//     aclass_quote // asset class of quote component
//     quote // asset id of quote component
//     lot // volume lot size
//     pair_decimals // scaling decimal places for pair
//     lot_decimals // scaling decimal places for volume
//     lot_multiplier // amount to multiply lot volume by to get currency volume
//     leverage_buy // array of leverage amounts available when buying
//     leverage_sell // array of leverage amounts available when selling
//     fees // fee schedule array in [volume, percent fee] tuples
//     fees_maker // maker fee schedule array in [volume, percent fee] tuples (if on maker/taker)
//     fee_volume_currency // volume discount currency
//     margin_call // margin call level
//     margin_stop // stop-out/liquidation margin level
// }


func main() {
  fmt.Println("\nGetting XXBT")
  xxbt := getAssetInfo("XXBT")
  fmt.Println(xxbt)
  // fmt.Println("\nGetting XXBT, ZUSD")
  // getAssetInfo("XXBT", "ZUSD")
  fmt.Println("\nGetting []string... of XXBT, ZUSD, XETH")
  XBTUSDETH := getAssetInfo([]string{"XXBT", "ZUSD", "XETH"}...)
  fmt.Println(XBTUSDETH)
  // fmt.Println("\nGetting XXBT, ZUSD, XETH")
  // getAssetInfo("XXBT", "ZUSD", "XETH")
  // fmt.Println("\nGetting All Assets")
  // getAssetInfo()

}

func (a asset) String() string {
  return fmt.Sprintf("Asset Class: %v\nAlternate Name: %v\nDecimals: %v\nDisplay Decimals: %v", a.Aclass, a.Altname, a.Decimals, a.DisplayDecimals)
}

func (a assetInfo) String() string {
  var assets string
  for assetName, asset := range(*a.Result) {
    assets += fmt.Sprintf("%v :\n%v\n", assetName, asset)
  }
  if len(a.Error) > 0 {
    return fmt.Sprintf("Error: %v\n%v", a.Error, assets)
  } else {
    return fmt.Sprintf("%v", assets)
  }
}


func getAssetInfo(assets ...string) assetInfo {
  url := "https://api.kraken.com/0/public/Assets"
  if len(assets) > 0 {
    url += "?asset="
    url += strings.Join(assets, ",")
  }
  resp, err := http.Get(url)
  if err != nil {
    panic(err)
  }
  assetinfo := assetInfo{}
  defer resp.Body.Close()
  err = json.NewDecoder(resp.Body).Decode(&assetinfo)
  if err != nil {
    panic(err)
  }
  return assetinfo
  // fmt.Println(resp.Body)
  // fmt.Println(assetinfo)
  // fmt.Println(*assetinfo.Result)
}


func getServerTime() {
  resp, err := http.Get("https://api.kraken.com/0/public/Time")
  if err != nil {
    panic(err)
  }
  st := serverTime{}
  defer resp.Body.Close()
  err = json.NewDecoder(resp.Body).Decode(&st)
  if err != nil {
    panic(err)
  }
  fmt.Println(resp.Body)
  fmt.Println(st)
  fmt.Println(*st.Result)
}
