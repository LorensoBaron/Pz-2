package main
import "fmt"

type Product struct {
    Name     string
    Category string
    Price    float64
}

func filterProducts(products []Product, maxPrice float64, category string) []Product {
    var filtered []Product
    for _, p := range products {
        if p.Price <= maxPrice && p.Category == category {
            filtered = append(filtered, p)
        }
    }
    return filtered
}

func main() {
    
    products := []Product{
        {"Товар1", "Электроника", 100},
        {"Товар2", "Электроника", 80},
        {"Товар3", "Одежда", 50},
        {"Товар4", "Электроника", 120},
    }
    
    result := filterProducts(products, 100, "Электроника")
    fmt.Println(result)
}
