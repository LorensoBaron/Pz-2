package main

import "fmt"

// 1. Объявление типа Order
type Order struct {
 ID          int
 Items       []int
 Total       float64
 Address     string
 IsCompleted bool
}

// 2. Создание карты (map)
var orders = make(map[int]Order)

// 3. Функция для добавления нового заказа в карту
func addOrder(order Order) {
 orders[order.ID] = order
}

func main() {
 // Пример использования

 // Создаем несколько заказов
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

 // Добавляем заказы в карту
 addOrder(order1)
 addOrder(order2)

 // Выводим информацию о заказах (для проверки)
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
