package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	// Каналы связи
	toSquare := make(chan int, 10)
	squared := make(chan int, 10)

	var wg sync.WaitGroup

	// 1-я горутина: генератор чисел (0-100)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(toSquare)

		rng := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 123))
		for i := 0; i < 10; i++ {
			num := int(rng.Uint64N(101))
			toSquare <- num
		}
	}()

	// 2-я горутина: возведение в квадрат
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(squared)

		for num := range toSquare {
			square := num * num
			squared <- square
		}
	}()

	wg.Wait()

	// Main: получаем все 10 квадратов
	fmt.Println("Финальные квадраты:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", <-squared)
	}
	fmt.Println()
}
