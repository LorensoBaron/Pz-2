package main

import (
    "fmt"
)

type Movie struct {
    Title   string
    Year    int
    Rating  float64
    Genres  []string
}

func main() {
    
    movies := []Movie{
        {"Интерстеллар", 2014, 8.6, []string{"Драма", "Фантастика"}},
        {"Тихоокеанский рубеж", 2013, 7.2, []string{"Экшн", "Фантастика"}},
        {"Матрица", 1999, 8.7, []string{"Научная фантастика", "Экшн"}},
        {"Гладиатор", 2000, 8.5, []string{"Исторический", "Драма"}},
        {"Живая сталь", 2011, 7.2, []string{"Экшн", "Спорт", "Наука-фэнтези"}},
    }

    bestRated := findBestRated(movies)
    fmt.Printf("Лучший фильм по рейтингу: %s\n", bestRated.Title)

    findByGenre(movies, "Драма")
}

func findBestRated(movies []Movie) *Movie {
    var topMovie *Movie = nil
    for _, m := range movies {
        if topMovie == nil || m.Rating > topMovie.Rating {
            topMovie = &m
        }
    }
    return topMovie
}

func findByGenre(movies []Movie, genre string) {
    fmt.Printf("\nФильмы жанра '%s':\n", genre)
    for _, m := range movies {
        for _, g := range m.Genres {
            if g == genre {
                fmt.Printf("- %s (%d)\n", m.Title, m.Year)
            }
        }
    }
}
