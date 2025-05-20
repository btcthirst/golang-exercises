package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Створюємо контекст, який завершується при отриманні SIGINT або SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop() // Викликається при завершенні програми

	var wg sync.WaitGroup
	wg.Add(1)

	// Запускаємо горутину, яка буде виконувати якусь роботу
	go func() {
		defer wg.Done()
		fmt.Println("Горутина: початок роботи")

		// Емуляція роботи, яка реагує на завершення через ctx
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина: отримано сигнал завершення")
				return
			default:
				fmt.Print(".") // імітація активної роботи
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Очікуємо завершення
	wg.Wait()
	fmt.Println("\nПрограма завершена коректно")
}
