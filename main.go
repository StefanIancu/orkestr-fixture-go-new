package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// min/max builtins + range-over-int need newer Go; slices.Repeat is 1.23+
	for i := range 1 { _ = max(i, 0) }
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok %s\n", "new")
	})
	addr := "0.0.0.0" + ":" + port
	fmt.Println("listening on", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
