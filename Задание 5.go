package main

import (
 "fmt"
 "strings"
)

func validateUser(name string, age int, email string) error {
 if name == "" || len(name) >= 50 {
  return fmt.Errorf("некорректное имя")
 }
 if age < 18 || age > 120 {
  return fmt.Errorf("некорректный возраст")
 }
 if !strings.Contains(email, "@") {
  return fmt.Errorf("некорректный email")
 }
 return nil
}

func main() {
 err := validateUser("Иван", 25, "ivan@example.com")
 if err != nil {
  fmt.Println("Ошибка:", err)
 } else {
  fmt.Println("Данные валидны")
 }
}





