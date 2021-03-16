/* ------------ License ------------ */
// BSD 3-Clause License
//
// Copyright (c) 2021, Seongmin Kim
// All rights reserved.
/* --------------------------------- */

package main

import (
	"github.com/shieldnet/gobit-sample/strategy"
	"github.com/shieldnet/gobit/jwtmaker"
	"sync"
)
const (
	Tfuel = "KRW-TFUEL"
)

func main() {
	coinList := []string{Tfuel}
	var strategies []*strategy.Strategy

	key := jwtmaker.Keys{
		Access: "my_access_key",
		Secret: []byte("my_secret_key"),
	}

	for _, coin := range coinList {
		s := &strategy.Strategy{
			Market:        coin,
			BuyCandleNum:  5,
			SellCandleNum: 5,
			QuitRate:      2.0,
			CandleUnit:    5,
			NextState:     "Init",
			Balance:       "0",
			TotalPrice:    "50000",
			Key: key,
		}
		strategies = append(strategies, s)
	}

	for true {
		wait := new(sync.WaitGroup)
		wait.Add(len(strategies))
		for _, s := range strategies {
			go s.Execute(wait)
		}
		wait.Wait()
	}
}
