package main

import "fmt"

const (
    Single = "single"
    Double = "double"
    Suite  = "suite"
)

const (
    Free        = "free"
    Booked      = "booked"
    Maintenance = "maintenance"
)

type HotelRoom struct {
    Type     string
    Status   string
    Price    float64
}

var hotel = map[string]HotelRoom{
    "101": {Type: Single, Status: Free, Price: 50},
    "102": {Type: Double, Status: Free, Price: 80},
    "201": {Type: Suite, Status: Free, Price: 150},
}

func bookRoom(roomNumber string) {
    room, exists := hotel[roomNumber]
    if !exists {
        fmt.Println("Комната не найдена")
        return
    }
    if room.Status != Free {
        fmt.Println("Комната недоступна для бронирования")
        return
    }
    room.Status = Booked
    hotel[roomNumber] = room
    fmt.Printf("Комната %s забронирована\n", roomNumber)
}

func main() {
    bookRoom("101")
    fmt.Println(hotel["101"])
}
