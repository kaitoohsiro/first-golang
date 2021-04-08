package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	// ルーティングを設定する
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})
	var port = "8080"
	log.Printf("Server listening on http://localhost:%s/", port)
	log.Print(http.ListenAndServe(":" + port, nil))
	// ポート8080番でサーバーを起動する
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		panic(err)
	}
}