package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yourusername/currency-converter/converter"
)

func main() {
	amount := flag.Float64("amount", 0, "Сумма для конвертации")
	from := flag.String("from", "USD", "Исходная валюта (например, USD)")
	to := flag.String("to", "EUR", "Целевая валюта (например, EUR)")
	flag.Parse()

	if *amount <= 0 {
		fmt.Println("Ошибка: укажите положительную сумму с помощью -amount")
		flag.Usage()
		os.Exit(1)
	}

	rates, err := converter.FetchRates(*from)
	if err != nil {
		log.Fatalf("Не удалось получить курсы: %v", err)
	}

	result, err := converter.Convert(*amount, *from, *to, rates)
	if err != nil {
		log.Fatalf("Ошибка конвертации: %v", err)
	}

	fmt.Printf("%.2f %s = %.2f %s\n", *amount, *from, result, *to)
}
