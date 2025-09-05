package main

import "fmt"

func main() {
 candidates := []string{"Анна", "Борис", "Виктор"}
 votes := []string{"Анна", "Борис", "Виктор", "Анна", "Анна", "Борис", "Виктор", "Анна", "Борис", "Борис", "Виктор", "Виктор", "Анна"}

 countVotes(candidates, votes)
}

func countVotes(candidates []string, votes []string) {
 voteCounts := make(map[string]int)
 totalVotes := len(votes)

 for _, vote := range votes {
  voteCounts[vote]++ // Инкремент, даже если нет ключа (по умолчанию 0)
 }

 fmt.Println("Результаты голосования:")
 for _, candidate := range candidates {
  count := voteCounts[candidate]
  percentage := float64(count) / float64(totalVotes) * 100
  fmt.Printf("%s: %d голосов (%.2f%%)\n", candidate, count, percentage)
 }
}

