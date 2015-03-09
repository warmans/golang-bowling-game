package main

type Game struct {
	rolls       [21]int
	currentRoll int
}

func (g *Game) Roll(pins int) {
	g.rolls[g.currentRoll] = pins
	g.currentRoll++
}

func (g *Game) GetScore() int {
	total := 0
	rollIndex := 0
	for frame := 0; frame < 10; frame++ {
		if g.isStrike(rollIndex) {
			total += 10 + g.strikeBonus(rollIndex)
			rollIndex++
		} else if g.isSpare(rollIndex) {
			total += 10 + g.spareBonus(rollIndex)
			rollIndex += 2
		} else {
			total += g.frameScore(rollIndex)
			rollIndex += 2
		}
	}
	return total
}

func (g *Game) isStrike(rollIndex int) bool {
	return (g.rolls[rollIndex] == 10)
}

func (g *Game) isSpare(rollIndex int) bool {
	return (g.rolls[rollIndex]+g.rolls[rollIndex+1] == 10)
}

func (g *Game) spareBonus(rollIndex int) int {
	return g.rolls[rollIndex+2]
}

func (g *Game) strikeBonus(rollIndex int) int {
	return g.rolls[rollIndex+1] + g.rolls[rollIndex+2]
}

func (g *Game) frameScore(rollIndex int) int {
	return g.rolls[rollIndex] + g.rolls[rollIndex+1]
}
