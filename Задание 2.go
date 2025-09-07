package main
import(
	"fmt"
)

func main() {
    var mainBaggage, handLuggage, extraLuggage float64

    fmt.Scanln(&mainBaggage)

    fmt.Scanln(&handLuggage)

    fmt.Scanln(&extraLuggage)

    total := mainBaggage + handLuggage + extraLuggage
    fmt.Printf("Общий вес багажа: %.2f кг\n", total)

}
