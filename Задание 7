package main

import "fmt"

type Employee struct {
 ID       int
 Name     string
 Position string
 Salary   float64
}

func calculatePayroll(employees []Employee) (float64, float64) {
 totalSalary := 0.0
 for _, emp := range employees {
  totalSalary += emp.Salary
 }

 if len(employees) == 0 {
  return 0, 0 
 }

 return totalSalary, totalSalary / float64(len(employees))
}

func main() {
 employees := []Employee{
  {ID: 1, Name: "Иван", Position: "Разраб", Salary: 100000},
  {ID: 2, Name: "Мария", Position: "Тестировщик", Salary: 80000},
 }

 total, avg := calculatePayroll(employees)
 fmt.Printf("Общий фонд: %.2f, Средняя: %.2f\n", total, avg)
}
