package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/mingodev/DND5eSheets/internal/pkg/dice"
)

func main() {

	defaultHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello World!\n")
		testRoll1 := dice.COIN.Roll(dice.DieRollParameters{})
		testRoll2 := dice.D20.Roll(dice.DieRollParameters{})
		io.WriteString(w, "Dice roll result 1 was: \n")
		io.WriteString(w, strconv.Itoa(testRoll1))
		io.WriteString(w, "\n")
		io.WriteString(w, "Dice roll result 1 was: \n")
		io.WriteString(w, strconv.Itoa(testRoll2))
		io.WriteString(w, "\n")
	}

	http.HandleFunc("/", defaultHandler)

	log.Println("Listening for requests at http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
