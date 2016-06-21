package tennis

import (
	"testing"
)

const (
	Love = "Love"
	Fifteen = "Fifteen"
	Thirty = "Thirty"
	Forty = "Forty"
)

func generateExpected(player1score, player2score string) string {
	return player1score + "-" + player2score
}

func TestZeroPointsFromEachSidePrintsLove(t *testing.T) {
	expected := generateExpected(Love, Love)

	score := getScore()
	if score != expected {
		t.Error("Error! Expected %s ", expected, " but got %s ", score)
	}
}

func TestLoveFifteen(t *testing.T) {
	expected := generateExpected(Love, Fifteen)
	wonPoint("player2")

	score := getScore()
	if score != expected {
		t.Error("Error! Expected %s ", expected, " but got %s ", score)
	}
}

func TestLoveThirty(t *testing.T) {
	expected := generateExpected(Love, Thirty)
	wonPoint("player2")
	wonPoint("player2")

	score := getScore()
	if score != expected {
		t.Error("Error! Expected %s ", expected, " but got %s ", score)
	}
}

func TestLoveForty(t *testing.T) {
	expected := generateExpected(Love, Forty)
	for i := 0; i < 3; i++ {
		wonPoint("player2")
	}

	score := getScore()
	if score != expected {
		t.Error("Error! Expected %s ", expected, " but got %s ", score)
	}
}

func TestFifteenForty(t *testing.T) {
	expected := generateExpected(Fifteen, Forty)
	wonPoint("player1")
	for i := 0; i < 3; i++ {
		wonPoint("player2")
	}

	score := getScore()
	if score != expected {
		t.Error("Error! Expected %s ", expected, " but got %s ", score)
	}
}

func TestWinForPlayer1(t *testing.T) {
	expected := "Win for player 1"
	for i := 0; i < 4; i++ {
		wonPoint("player1")
	}
	wonPoint("player2")
	wonPoint("player2")

	score := getScore()
	if score != expected {
		t.Error("Error! Expected %s ", expected, " but got %s ", score)
	}
}

func TestWinForPlayer2(t *testing.T) {
	expected := "Win for player 2"

	wonPoint("player2")
	wonPoint("player2")
	wonPoint("player2")
	wonPoint("player2")
	wonPoint("player1")
	wonPoint("player1")

	score := getScore()
	if score != expected {
		t.Error("Error! Expected %s ", expected, " but got %s ", score)
	}
}

func TestDeuce(t *testing.T) {
	expected := "Deuce"

	wonPoint("player2")
	wonPoint("player2")
	wonPoint("player2")
	wonPoint("player1")
	wonPoint("player1")
	wonPoint("player1")

	score := getScore()
	if score != expected {
		t.Error("Error! Expected %s ", expected, " but got %s ", score)
	}
}

func TestAdvantagePlayer1(t *testing.T) {
	expected := "Advantage player1"

	wonPoint("player2")
	wonPoint("player2")
	wonPoint("player2")
	wonPoint("player1")
	wonPoint("player1")
	wonPoint("player1")
	wonPoint("player1")

	score := getScore()
	if score != expected {
		t.Error("Error! Expected %s ", expected, " but got %s ", score)
	}
}

func TestAdvantagePlayer2(t *testing.T) {
	expected := "Advantage player2"

	wonPoint("player2")
	wonPoint("player2")
	wonPoint("player2")
	wonPoint("player2")
	wonPoint("player1")
	wonPoint("player1")
	wonPoint("player1")

	score := getScore()
	if score != expected {
		t.Error("Error! Expected %s ", expected, " but got %s ", score)
	}
}

