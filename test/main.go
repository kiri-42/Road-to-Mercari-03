package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const DAI_KICHI_I = 0

type Task struct {
	Fortune string    `json:"fortune"`
	Date    time.Time `json:"date"`
}

var tasks = []Task{{
	Fortune: "Dai-kichi",
	Date:    time.Now(),
}, {
	Fortune: "kichi",
	Date:    time.Now(),
}, {
	Fortune: "Chuu-kichi",
	Date:    time.Now(),
}, {
	Fortune: "Sho-kichi",
	Date:    time.Now(),
}, {
	Fortune: "Tue-kichi",
	Date:    time.Now(),
}, {
	Fortune: "Kyo",
	Date:    time.Now(),
}, {
	Fortune: "Dai-kyo",
	Date:    time.Now(),
}}

func main() {
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":4242", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&tasks[getIndex()]); err != nil {
		log.Fatal(err)
	}

	_, err := fmt.Fprint(w, buf.String())
	if err != nil {
		return
	}
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
