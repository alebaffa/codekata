package bowling

import "testing"

func TestTwoRolls(t *testing.T) {
	game := Game{}
	game.Roll(1)
	game.Roll(2)
	score := game.GetScore()
	printResult(score, 3, t)

}

func TestSpare(t *testing.T) {
	game := Game{}
	game.Roll(5)
	game.Roll(5)
	game.Roll(3)
	score := game.GetScore()
	printResult(score, 16, t)

}

func TestSimpleFrameAfterSpare(t *testing.T) {
	game := Game{}
	game.Roll(7)
	game.Roll(3)
	game.Roll(3)
	game.Roll(2)
	score := game.GetScore()
	printResult(score, 18, t)

}

func TestStrike(t *testing.T) {
	game := Game{}
	game.Roll(10)
	game.Roll(3)
	game.Roll(6)
	score := game.GetScore()
	printResult(score, 28, t)

}

func TestPerfectGame(t *testing.T) {
	game := Game{}
	i := 0
	for i < 12 {
		game.Roll(10)
		i++
	}
	score := game.GetScore()
	printResult(score, 300, t)

}

func printResult(actualScore int, expectedScore int, t *testing.T) {
	if actualScore != expectedScore {
		t.Error("Got score ", actualScore, ", but expected ", expectedScore)
	}
}
