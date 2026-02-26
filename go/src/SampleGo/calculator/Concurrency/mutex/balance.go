package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	var bankBalance int
	income := []Income{
		{Source: "main job", Amount: 1000},
		{Source: "investment", Amount: 200},
		{Source: "pocket money", Amount: 10},
	}

	fmt.Println("initial income is  ", bankBalance)

	wg.Add(len(income))

	for i, in := range income {

		go func(i int, income Income) {
			defer wg.Done()
			for i := 0; i < 52; i++ {
				mutex.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp

				fmt.Printf("on week %d you earned %d from %s\n", i, bankBalance, income.Source)
				mutex.Unlock()
			}
		}(i, in)
	}

	wg.Wait()
	fmt.Println("Final balance is  ", bankBalance)
}
