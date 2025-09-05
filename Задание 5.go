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
 err := validateUser("Иван Иванов", 30, "ivan@example.com")
 if err != nil {
  fmt.Println("Ошибка:", err)
 }

 err = validateUser("", 25, "test@example.com")
 if err != nil {
  fmt.Println("Ошибка:", err)
 }

 err = validateUser("Петр Петров", 15, "petr@example.com")
 if err != nil {
  fmt.Println("Ошибка:", err)
 }

 err = validateUser("Сергей Сергеев", 40, "invalid-email")
 if err != nil {
  fmt.Println("Ошибка:", err)
 }
}


