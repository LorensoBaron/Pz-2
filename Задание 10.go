package main
import (
"fmt"
"strings"
)

type TextStatistics struct {
   SymbolsCount int
   WordsCount int
   SentencesCount int
}

func textStats(text string) *TextStatistics { 
    stats := &TextStatistics{}
    
   stats.SymbolsCount = len([]rune(text))
   
   words := strings.Fields(text)
   
   stats.WordsCount = len(words)
   
  sentenceDelimiters := []string{".", "!", "?"} 
  for _, delimiter := range sentenceDelimiters { 
      count := strings.Count(text, delimiter) 
      
      stats.SentencesCount += count
  }
   return stats
}
func main() {
text := "Здраствуйте! Добрый вечер. Как дела?"
result := textStats(text)
fmt.Printf("Символов: %d ", result.SymbolsCount)
fmt.Printf("Слов: %d ", result.WordsCount)
fmt.Printf("Предложений: %d ", result.SentencesCount)
}
