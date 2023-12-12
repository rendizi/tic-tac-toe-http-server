package main

import (
	"fmt"
	"github.com/rendizi/tic-tac-toe-http-server/http/server/handler"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/newgame", handler.NewGame)
	http.HandleFunc("/game", handler.Game)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		if err == http.ErrServerClosed {
			fmt.Println("server closed")
		} else {
			fmt.Printf("error starting server: %s\n", err)
			os.Exit(1)
		}
	}
}
