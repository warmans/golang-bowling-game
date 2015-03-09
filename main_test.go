package main

import "testing"

func RollMany(game *Game, n int, pins int) {
	for i := 0; i < n; i++ {
		game.Roll(pins)
	}
}

func RollSpare(game *Game) {
	game.Roll(5)
	game.Roll(5)
}

func RollStrike(game *Game) {
	game.Roll(10)
}

func TestGutterGame(t *testing.T) {
	game := Game{}

	RollMany(&game, 20, 0)

	if game.GetScore() != 0 {
		t.Error("Score was incorrect", 0, game.GetScore())
	}
}

func TestAllOnes(t *testing.T) {
	game := Game{}

	RollMany(&game, 20, 1)

	if game.GetScore() != 20 {
		t.Error("Score was incorrect", 20, game.GetScore())
	}
}

func TestOneSpare(t *testing.T) {
	game := Game{}
	RollSpare(&game)
	game.Roll(3)

	RollMany(&game, 17, 0)

	if game.GetScore() != 16 {
		t.Error("Score was incorrect", 16, game.GetScore())
	}
}

func TestOneStrike(t *testing.T) {
	game := Game{}

	RollStrike(&game)
	game.Roll(3)
	game.Roll(4)

	RollMany(&game, 16, 0)

	if game.GetScore() != 24 {
		t.Error("Score was incorrect", 24, game.GetScore())
	}
}

func testPerfectGame(t *testing.T) {
	game := Game{}
	RollMany(&game, 12, 10)
	if game.GetScore() != 300 {
		t.Error("Score was incorrect", 300, game.GetScore())
	}
}
