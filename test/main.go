package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Task struct {
	Fortune string     `json:"fortune"`
	DueDate time.Time `json:"due_date"`
}

var tasks = []Task{{
	Fortune: "Dai-kichi",
	DueDate: time.Now(),
}, {
	Fortune: "kichi",
	DueDate: time.Now(),
}, {
	Fortune: "Chuu-kichi",
	DueDate: time.Now(),
}, {
	Fortune: "Sho-kichi",
	DueDate: time.Now(),
}, {
	Fortune: "Sue-kichi",
	DueDate: time.Now(),
}, {
	Fortune: "Kyo",
	DueDate: time.Now(),
}, {
	Fortune: "Dai-kyo",
	DueDate: time.Now(),
}}

func main() {
	handler1 := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		if err := enc.Encode(&tasks); err != nil {
			log.Fatal(err)
		}
		fmt.Println(buf.String())

		_, err := fmt.Fprint(w, buf.String())
		if err != nil {
			return
		}
	}

	// GET /tasks
	http.HandleFunc("/", handler1)
	log.Fatal(http.ListenAndServe(":4242", nil))
}
