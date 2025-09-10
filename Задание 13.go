package main

import "fmt"

func main() {
    expenses := make(map[string]float64)

    expenses["Еда"] = 15000
    expenses["Транспорт"] = 5000
    expenses["Развлечения"] = 3000

    expenses["Еда"] += 2000

    totalExpenses := 0.0
    for _, amount := range expenses {
        totalExpenses += amount
    }

    fmt.Println("Итоговые расходы:", totalExpenses)
}
