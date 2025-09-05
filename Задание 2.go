package main
import(
	"fmt"
)

func main() {
    var mainBaggage, handLuggage, extraLuggage float64
// основной багаж
    fmt.Scanln(&mainBaggage)
// ручная кладь
    fmt.Scanln(&handLuggage)
// доп ручная кладь
    fmt.Scanln(&extraLuggage)

    total := mainBaggage + handLuggage + extraLuggage
    fmt.Printf("Общий вес багажа: %.2f кг\n", total)
}