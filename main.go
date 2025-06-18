package main

import (
    "database/sql"
    "html/template"
    "log"
    "net/http"

    _ "github.com/mattn/go-sqlite3"
)

type Verse struct {
    Book    string
    Chapter int
    Verse   int
    Text    string
}

func main() {
    http.HandleFunc("/", serveRandomChapter)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    log.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveRandomChapter(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("sqlite3", "./bible/ARC.sqlite")
    if err != nil {
        http.Error(w, "Database connection error", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    query := `
    WITH selected_chapter AS (
      SELECT DISTINCT
             b.id,
             v.chapter
        FROM book AS b
        JOIN verse AS v
          ON b.id = v.book_id
    ORDER BY RANDOM()
       LIMIT 1
    )
       SELECT b.name AS book,
              v.chapter,
              v.verse,
              v.text
         FROM verse AS v
         JOIN book AS b
           ON v.book_id = b.id
         JOIN selected_chapter
           ON b.id = selected_chapter.id
          AND v.chapter = selected_chapter.chapter
     ORDER BY v.id;
    `

    rows, err := db.Query(query)
    if err != nil {
        http.Error(w, "Query execution error", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var verses []Verse
    for rows.Next() {
        var verse Verse
        if err := rows.Scan(&verse.Book, &verse.Chapter, &verse.Verse, &verse.Text); err != nil {
            http.Error(w, "Error scanning row", http.StatusInternalServerError)
            return
        }
        verses = append(verses, verse)
    }

    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, "Template parsing error", http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, verses)
}
