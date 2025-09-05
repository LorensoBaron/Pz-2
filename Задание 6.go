package main

import "fmt"

func main() {
 postTags := [][]string{
  {"go", "backend"},
  {"git", "go", "tools"},
  {"backend", "docker"},
 }

 uniqueTags := getUniqueTags(postTags)

 fmt.Println("Уникальные теги:")
 for tag := range uniqueTags {
  fmt.Println(tag)
 }
}

func getUniqueTags(postTags [][]string) map[string]bool {
 uniqueTags := make(map[string]bool)
 for _, tags := range postTags {
  for _, tag := range tags {
   uniqueTags[tag] = true
  }
 }
 return uniqueTags
}
