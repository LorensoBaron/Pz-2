package main

import "fmt"

type Order struct {
 ID          int
 Items       []int
 Total       float64
 Address     string
 IsCompleted bool
}

var orders = make(map[int]Order)

func addOrder(order Order) {
 orders[order.ID] = order
}

func main() {
 order1 := Order{
  ID:          1,
  Items:       []int{101, 102, 103},
  Total:       150.75,
  Address:     "ул. Пушкина, д. 10",
  IsCompleted: false,
 }

 order2 := Order{
  ID:          2,
  Items:       []int{201, 202},
  Total:       75.50,
  Address:     "пр. Ленина, д. 25",
  IsCompleted: true,
 }

 addOrder(order1)
 addOrder(order2)

 fmt.Println("Информация о заказах:")
 for id, order := range orders {
  fmt.Printf("Заказ ID: %d\n", id)
  fmt.Printf("  Items: %v\n", order.Items)
  fmt.Printf("  Total: %.2f\n", order.Total)
  fmt.Printf("  Address: %s\n", order.Address)
  fmt.Printf("  IsCompleted: %t\n", order.IsCompleted)
  fmt.Println("----------------------")
 }
}

