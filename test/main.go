package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"math/rand"
)

const DAI_KICHI_I = 0

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
	Fortune: "Tue-kichi",
	DueDate: time.Now(),
}, {
	Fortune: "Kyo",
	DueDate: time.Now(),
}, {
	Fortune: "Dai-kyo",
	DueDate: time.Now(),
}}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		if err := enc.Encode(&tasks[getIndex()]); err != nil {
			log.Fatal(err)
		}
		fmt.Println(buf.String())

		_, err := fmt.Fprint(w, buf.String())
		if err != nil {
			return
		}
	}

	// GET /tasks
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":4242", nil))
}

func getIndex() (i int) {
	rand.Seed(time.Now().UnixNano())
	i = rand.Intn(len(tasks))

	now := time.Now()
	if isNewYear(now) {
		fmt.Println("isNewYear")
		i = DAI_KICHI_I
	}
	return i
}

func isNewYear(t time.Time) bool {
	day := t.Day()
	if int(t.Month()) == 1 && 1 <= day && day <= 3 {
		return true
	}
	return false
}
