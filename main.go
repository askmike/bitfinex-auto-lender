package main

import (
  "time"
  "strings"
  "log"
  "github.com/eAndrius/bitfinex-go"
  "github.com/spf13/viper"
)


type config struct {
  Live bool
  CheckEvery float64
  LendDays int
  MinimumUSD float64
  Currency string
  ApiKey string
  ApiSecret string
}

var C config
var api *bitfinex.API

func main() {

  viper.SetConfigName("config")
  viper.AddConfigPath(".")
  viper.ReadInConfig()

  err := viper.Unmarshal(&C)
  if err != nil {
    log.Println("unable to decode into struct, %v", err)
  }

  api = bitfinex.New(C.ApiKey, C.ApiSecret)
  t := time.NewTicker(time.Minute * time.Duration(C.CheckEvery))
  for {
      lend()
      <-t.C
  }
}

func cancelOrders() {
  if C.Live != true {
    log.Print("Not live, not canceling orders.")
    return
  }

  err := api.CancelActiveOffersByCurrency(C.Currency)
  if err != nil {
    log.Print("Unable to cancel orders: " + err.Error())
  }
}

func lend() {

  cancelOrders()

  log.Print("Checking balance..")
  balance, err := getBalance()
  if err != nil {
    return
  }

  if balance == 0 {
    log.Print("Empty balance, skipping.")
    return
  }

  log.Print("Balance: ", balance)

  log.Print("Checking minimum..")
  minimum, err := getMinimum()
  if err != nil {
    return
  }
  log.Print("Minumum: ", minimum)

  if(balance < minimum) {
    log.Print("Balance lower than minimum lending amount, skipping.") 
    return
  }
  
  log.Print("Checking best rate offered..")
  topAsk, err := getTopAsk();
  if err != nil {
    return
  }

  log.Print("Best rate: ", topAsk)
  log.Print("Creating offer: ", balance, "@", topAsk/365, " (", C.LendDays, " days)")

  if C.Live == false {
    log.Println("Not live, not placing offer.")
    return
  }

  _, err = api.NewOffer(strings.ToUpper(C.Currency), balance, topAsk, C.LendDays, "lend")
  if err != nil {
    log.Println("Failed to place new offer: " + err.Error())
  }
}



// helpers

func getMinimum() (float64, error) {
  if(C.Currency == "usd") {
    return C.MinimumUSD, nil
  }

  ticker, err := api.Ticker(C.Currency + "usd")
  if err != nil {
    log.Print("Error getting ticker: " + err.Error())
    return 0, err
  }

  minimum := C.MinimumUSD / ticker.Mid

  return minimum, nil
}

func getBalance() (float64, error) {
  balance, err := api.WalletBalances()
  if err != nil {
    log.Print("Error getting balance: " + err.Error())
    return 0, err
  }

  depositBalance := balance[bitfinex.WalletKey{"deposit", C.Currency}]

  return depositBalance.Available, nil
}

func getTopAsk() (float64, error) {
  lendbook, err := api.Lendbook(C.Currency, 0, 1)
  if err != nil {
    log.Println("Error getting lendbook: " + err.Error())
    return 0, err 
  }

  topAsk := lendbook.Asks[0].Rate

  return topAsk, nil
}