package main

import (
    "fmt"
    "time"
)

// Сенсорные данные
type SensorData struct {
    SensorID   string    
    Temperature float64  
    Humidity    float64  
    Timestamp   time.Time 
}

func main() {
    
    data := make([]SensorData, 0, 6)

    now := time.Now()
    data = append(data,
        SensorData{"Sensor_1", 21.5, 55.0, now},
        SensorData{"Sensor_2", 22.1, 58.0, now.Add(time.Hour)},
        SensorData{"Sensor_1", 20.8, 54.5, now.Add(2 * time.Hour)},
        SensorData{"Sensor_3", 23.0, 60.0, now.Add(3 * time.Hour)},
        SensorData{"Sensor_2", 21.9, 57.0, now.Add(4 * time.Hour)})

    avgTemp := averageTemperature(data)
    fmt.Printf("Средняя температура за сутки: %.2f°C\n", avgTemp)
}

func averageTemperature(data []SensorData) float64 {
    sum := 0.0
    for _, d := range data {
        sum += d.Temperature
    }
    return sum / float64(len(data))
}
