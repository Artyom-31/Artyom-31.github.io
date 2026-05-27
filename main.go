package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// В Go 1.20+ работает БЕЗ Seed автоматически!
	fmt.Println(rand.Intn(100)) // Разные числа при каждом запуске
}
