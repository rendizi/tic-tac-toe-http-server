package handler

import (
	"encoding/json"
	"fmt"
	"github.com/rendizi/tic-tac-toe-http-server/pkg/game"
	"net/http"
	"strconv"
)

var (
	Games map[int]*game.Net
)

func init() {
	Games = make(map[int]*game.Net)
}
func NewGame(w http.ResponseWriter, r *http.Request) {
	Net, err := game.NewNet()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Games[Net.Id] = Net

	fmt.Fprintf(w, "Your unique game id: %d. You can join it via /game", Net.Id)
}

func Game(w http.ResponseWriter, r *http.Request) {
	gameIDHeader := r.Header.Get("ID")
	if gameIDHeader == "" {
		http.Error(w, "Game ID not provided in header", http.StatusBadRequest)
		return
	}

	gameID, err := strconv.Atoi(gameIDHeader)
	if err != nil {
		http.Error(w, "Invalid Game ID provided in header", http.StatusBadRequest)
		return
	}

	Net, ok := Games[gameID]
	if !ok {
		http.Error(w, "Game ID not found", http.StatusNotFound)
		return
	}
	if r.Body == nil {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	var p struct {
		X       string `json:"x"`
		Y       string `json:"y"`
		IsFirst bool   `json:"isFirst"`
	}
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	x, err := strconv.Atoi(p.X)
	if err != nil {
		http.Error(w, "Invalid data for 'x'", http.StatusBadRequest)
		return
	}
	y, err := strconv.Atoi(p.Y)
	if err != nil {
		http.Error(w, "Invalid data for 'y'", http.StatusBadRequest)
		return
	}
	y--
	x--
	if x >= 3 || x < 0 || y >= 3 || y < 0 {
		http.Error(w, "Not valid data", http.StatusBadRequest)
		return
	}
	isWinner, err := Net.Set(2-y, x, p.IsFirst)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			cell := Net.Grid[i][j]
			if cell == "" {
				fmt.Fprintf(w, " ")
			} else {
				fmt.Fprintf(w, "%s", cell)
			}

			if j != 2 {
				fmt.Fprintf(w, "|")
			}
		}
		fmt.Fprintf(w, "\n")
		if i != 2 {
			fmt.Fprintf(w, "-+-+-\n")
		}
	}
	if isWinner {
		fmt.Fprintln(w, "\nGame over!")
		Net, err = game.NewNet()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Games[gameID] = Net
		return
	}
}
