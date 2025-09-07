package main
import(
	"fmt"
)

func main() {
    days := []int{ 2, 3, 4, 5, 6, 7, 4, 5} 

 weekdayPrice := 2100
 weekendPrice := 2850

 totalCost := 0
 for _, day := range days {
  if day >= 1 && day <= 4 { // Пн, Вт, Ср, Чт
   totalCost += weekdayPrice
  } else if day >= 5 && day <= 7 { // Пт, Сб, Вс
   totalCost += weekendPrice
  } else {
   fmt.Println("Ошибка: Некорректный день недели:", day) 
   return 
  }
 }

 fmt.Println(totalCost)	

}
