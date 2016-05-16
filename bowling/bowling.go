package bowling

type Game struct {
	throws       [21]int
	currentThrow int
}

const maxFrames = 10

func (game *Game) Roll(pins int) {
	game.throws[game.currentThrow] = pins
	game.currentThrow++
}

func (game *Game) GetScore() int {
	score := 0
	indexRoll := 0

	for frame := 0; frame < maxFrames; frame++ {
		if game.isSpare(indexRoll) {
			score += 10 + game.nextBallForSpare(indexRoll)
			indexRoll += 2

		} else if game.isStrike(indexRoll) {
			score += 10 + game.nextTwoBallsForStrike(indexRoll)
			indexRoll++

		} else {
			score += game.normalFrame(indexRoll)
			indexRoll += 2
		}
	}
	return score
}

func (game *Game) isSpare(indexRoll int) bool {
	return game.throws[indexRoll]+game.throws[indexRoll+1] == 10
}

func (game *Game) nextBallForSpare(indexRoll int) int {
	return game.throws[indexRoll+2]
}

func (game *Game) isStrike(indexRoll int) bool {
	return game.throws[indexRoll] == 10
}

func (game *Game) nextTwoBallsForStrike(indexRoll int) int {
	return game.throws[indexRoll+1] + game.throws[indexRoll+2]
}

func (game *Game) normalFrame(indexRoll int) int {
	return game.throws[indexRoll] + game.throws[indexRoll+1]
}
