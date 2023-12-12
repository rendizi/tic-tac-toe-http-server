package handler

func NewGame() {
	Net, err := tictac.NewNet()
	if err != nil {
		return err
	}
}
