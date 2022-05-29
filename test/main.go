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

type Omikuji struct {
	Fortune string `json:"fortune"`
	Massage string `json:"massage"`
}

var omikujiList = []Omikuji{{
	Fortune: "Dai-kichi",
	Massage: "課題が一発クリア！！",
}, {
	Fortune: "kichi",
	Massage: "レビュワーにコードを褒められる",
}, {
	Fortune: "Chuu-kichi",
	Massage: "マッチングしたレビューがNetPractice",
}, {
	Fortune: "Sho-kichi",
	Massage: "よくわからんAchievementでWalletゲット",
}, {
	Fortune: "Sue-kichi",
	Massage: "1週間後にレビューセールが",
}, {
	Fortune: "Kyo",
	Massage: "作業中のパソコンが固まる",
}, {
	Fortune: "Dai-kyo",
	Massage: "TIG",
}}

func main() {
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":4242", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&omikujiList[getIndex()]); err != nil {
		log.Fatal(err)
	}

	_, err := fmt.Fprint(w, buf.String())
	if err != nil {
		return
	}
}

func getIndex() (i int) {
	rand.Seed(time.Now().UnixNano())
	i = rand.Intn(len(omikujiList))

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
