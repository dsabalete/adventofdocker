package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Record new visit
		timestamp := time.Now().Format(time.RFC3339) + "\n"
		os.MkdirAll("/data", 0755)
		f, _ := os.OpenFile("/data/visits.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		f.WriteString(timestamp)
		f.Close()

		// Read and display all visits
		data, err := os.ReadFile("/data/visits.txt")
		if err != nil {
			fmt.Fprintf(w, "Error reading visits!")
			return
		}

		fmt.Fprintf(w, "New visit recorded at %s\n\nAll visits:\n%s", timestamp, string(data))
	})

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
