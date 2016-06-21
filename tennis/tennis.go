package tennis

var player1score = 0
var player2score = 0

var score = [4]string{"Love", "Fifteen", "Thirty", "Forty"}

func getScore() string {
	switch {
	case player1Wins():
		resetScores()
		return "Win for player 1"
	case player2Wins():
		resetScores()
		return "Win for player 2"
	case isDeuce():
		resetScores()
		return "Deuce"
	case isAdvantagePlayer1():
		resetScores()
		return "Advantage player1"
	case isAdvantagePlayer2():
		resetScores()
		return "Advantage player2"
	default:
		currentScore := score[player1score] + "-" + score[player2score]
		resetScores()
		return currentScore
	}
}

func player1Wins() bool {
	return player1score == 4 && (player1score == (player2score + 2))
}
func player2Wins() bool {
	return player2score == 4 && (player2score == (player1score + 2))
}

func isDeuce() bool {
	return player1score == 3 && player2score == 3
}

func isAdvantagePlayer1() bool {
	return player1score == 4 && player2score == 3
}

func isAdvantagePlayer2() bool {
	return player2score == 4 && player1score == 3
}

func wonPoint(player string) {
	switch player{
	case "player1":
		player1score++
	case "player2":
		player2score++
	}
}

func resetScores() {
	player1score = 0
	player2score = 0
}