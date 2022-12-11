package strategyguide

import "errors"

// Opponent is going to play:
// A for Rock,
// B for Paper,
// C for Scissors.
const opponentsRock = 'A'
const opponentsPaper = 'B'
const opponentsScissors = 'C'

// What you should play in response:
// X for Rock,
// Y for Paper,
// Z for Scissors.
const yourRock = 'X'
const yourPaper = 'Y'
const yourScissors = 'Z'

// The score for the outcome of the round:
// 0 if you lost
// 3 if the round was a draw
// 6 if you won.
const pointsIfYouLost = 0
const pointsIfDraw = 3
const pointsIfYouWon = 6

type RoundResult int

const (
	YouLost RoundResult = iota
	Draw                = 1
	YouWon              = 2
)

// The score for a single round is:
// the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) +
// the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).
func GetScoreRound(opponentsChoice, yourChoice rune) (int, error) {
	scoreShapeYouSelected := map[rune]int{
		yourRock:     1,
		yourPaper:    2,
		yourScissors: 3,
	}

	roundOutcomeScore, err := calculateRoundOutcomeScore(opponentsChoice, yourChoice)
	if err != nil {
		return 0, err
	}

	return scoreShapeYouSelected[yourChoice] + roundOutcomeScore, nil
}

func whoWonTheRound(opponentsChoice, yourChoice rune) (RoundResult, error) {
	if opponentsChoice == opponentsRock && yourChoice == yourRock {
		return Draw, nil
	}
	if opponentsChoice == opponentsRock && yourChoice == yourPaper {
		return YouWon, nil
	}
	if opponentsChoice == opponentsRock && yourChoice == yourScissors {
		return YouLost, nil
	}

	if opponentsChoice == opponentsPaper && yourChoice == yourRock {
		return YouLost, nil
	}
	if opponentsChoice == opponentsPaper && yourChoice == yourPaper {
		return Draw, nil
	}
	if opponentsChoice == opponentsPaper && yourChoice == yourScissors {
		return YouWon, nil
	}

	if opponentsChoice == opponentsScissors && yourChoice == yourRock {
		return YouWon, nil
	}
	if opponentsChoice == opponentsScissors && yourChoice == yourPaper {
		return YouLost, nil
	}
	if opponentsChoice == opponentsScissors && yourChoice == yourScissors {
		return Draw, nil
	}

	return RoundResult(0), errors.New("invalid round")
}

func calculateRoundOutcomeScore(opponentsChoice, yourChoice rune) (int, error) {
	roundResult, err := whoWonTheRound(opponentsChoice, yourChoice)
	if err != nil {
		return 0, err
	}

	if roundResult == YouLost {
		return pointsIfYouLost, nil
	}

	if roundResult == Draw {
		return pointsIfDraw, nil
	}

	if roundResult == YouWon {
		return pointsIfYouWon, nil
	}

	return 0, errors.New("invalid round result")
}