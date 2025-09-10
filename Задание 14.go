package main

import "fmt"

type InventoryItem struct {
    Name        string
    Weight      float64
    IsQuestItem bool
}

func main() {
    
    inventory := []InventoryItem{
        {Name: "Меч", Weight: 3.5, IsQuestItem: false},
        {Name: "Зелье здоровья", Weight: 0.5, IsQuestItem: true},
        {Name: "Щит", Weight: 4.0, IsQuestItem: false},
        {Name: "Кольцо силы", Weight: 0.2, IsQuestItem: true},
        {Name: "Арбалет", Weight: 2.5, IsQuestItem: false},
    }

    totalWeight := calculateTotalWeight(inventory)
    fmt.Printf("Общий вес инвентаря: %.1f кг\n", totalWeight)
}

func calculateTotalWeight(items []InventoryItem) float64 {
    total := 0.0
    for _, item := range items {
        total += item.Weight
    }
    return total
}
