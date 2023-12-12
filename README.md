# tic-tac-toe-http-server
The Tic Tac Toe HTTP server in Golang is a lightweight web application that implements the classic Tic Tac Toe game.

Get's started

There are two requests - for a new game and for the next step. Below is an example of a request for a new game. A unique game ID is returned through which you can play

```
curl -X POST http://127.0.0.1:3333/newgame
```

Below is an example request for the next step. Put your game ID, cell coordinates and what to put (isFirst = true -> cross)

```
curl -X POST \
  -H "ID: 123" \ # Replace "123" with the actual Game ID
  -H "Content-Type: application/json" \
  -d '{"x": "1", "y": "2", "isFirst": true}' \
  http://localhost/game
```

At the end of the game, you can not create a new game and send a request to the same ID, the matrix will update itself
