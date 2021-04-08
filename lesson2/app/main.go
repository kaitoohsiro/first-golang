package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	port := "8080"
	http.HandleFunc("/", handleIndex)
	log.Printf("Server listening on http://localhost:%s/", port)
	log.Print(http.ListenAndServe(":" + port, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("root/index.html")
	if err != nil {
		log.Fatalf("template error: %v", err)
	}
	if err := t.Execute(w, struct {
		Title   string
		Message string
		Time    time.Time
	}{
		Title:   "テストページ",
		Message: "こんにちは！",
		Time:    time.Now(),
	}); err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}