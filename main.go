package main

import (
	"cripto/api"
	"fmt"
	"sync"
)

func main() {
	currencies := []string {"BTC", "ETH", "SOL"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1)
		go func () {
			getCurrencyData(currency)
			wg.Done()
		}()
	}
	wg.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The rate for %s is %f\n", rate.Currency, rate.Price)
}
