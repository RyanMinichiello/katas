package war

func PlayGame(decks [][]int) int {
	game := struct {
		gameOver bool
		winner int
	}{
		false,
		-1,
	}
	playerOne := decks[0]
	playerTwo := decks[1]
	for game.gameOver != true {
		if len(playerOne) == 0 {
			game.gameOver = true
			game.winner = 0
			continue
		}

		if len(playerTwo) == 0 {
			game.gameOver = true
			game.winner = 1
			continue
		}
		onePlays := playerOne[0]
		twoPlays := playerTwo[0]
		if onePlays == twoPlays {
			// WAR
			decks := [][]int{playerOne[1:], playerTwo[1:]}
			results := War(decks)
			playerOne = append(playerOne, results[0]...)
			playerTwo = append(playerTwo, results[1]...)
			continue
		}
		if onePlays > twoPlays {
			playerOne = append(playerOne, twoPlays)
			playerTwo = playerTwo[1:]
			continue
		}
		if twoPlays > onePlays {
			playerTwo = append(playerTwo, onePlays)
			playerOne = playerOne[1:]
			continue
		}
	}
	return game.winner
}

func War(decks [][]int) [][]int {
	// TODO: init safe length arrays for War games
	results := [][]int{{1}, {2}}
	return results
}