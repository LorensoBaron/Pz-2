package main

import "fmt"

type LogEntry struct {
    IP        string
    HTTPCode  int
    Timestamp string
}

var logs = []LogEntry{
    {"192.169.1.1", 200, "2024-04-27 10:00:00"},
    {"192.168.1.2", 404, "2024-04-27 10:01:00"},
    {"192.168.1.3", 553, "2024-04-27 10:02:00"},
    {"192.168.1.4", 302, "2024-04-27 10:03:00"},
}

func filterErrors(entries []LogEntry) []string {
    var errors []string
    for _, entry := range entries {
        if entry.HTTPCode >= 400 && entry.HTTPCode < 500 {
            errors = append(errors, fmt.Sprintf("%v (4xx)", entry))
        } else if entry.HTTPCode >= 500 && entry.HTTPCode < 600 {
            errors = append(errors, fmt.Sprintf("%v (5xx)", entry))
        }
    }
    return errors
}

func main() {
   for _, log := range filterErrors(logs) {
    fmt.Println(log)
}
}
