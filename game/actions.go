package game

func handleActionQuit() {
	userQuit = true
}

func handleActionReset() {
	renderMutex.Lock()
	defer renderMutex.Unlock()

	generateMineField()
}

func handleActionInput(char rune) {
	renderMutex.Lock()
	defer renderMutex.Unlock()

}
