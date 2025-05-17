package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycles := 5

		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		if val := r.Form.Get("cycles"); val != "" {
			if conv, err := strconv.Atoi(val); err == nil {
				cycles = conv
			}
		}

		lissajous(w, cycles)
	})

	log.Println("Serving on http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
