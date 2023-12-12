package main

import (
	"fmt"
	"github.com/aivanov/game/http/server/handler"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/newgame", handler.NewGame)
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
