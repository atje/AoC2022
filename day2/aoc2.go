/*

--- Day 2: Rock Paper Scissors ---
The Elves begin to set up camp on the beach. To decide whose tent gets to be closest to the snack storage, a giant Rock Paper Scissors tournament is already in progress.

Rock Paper Scissors is a game between two players. Each game contains many rounds; in each round, the players each simultaneously choose one of Rock, Paper, or Scissors using a hand shape. Then, a winner for that round is selected: Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock. If both players choose the same shape, the round instead ends in a draw.

Appreciative of your help yesterday, one Elf gives you an encrypted strategy guide (your puzzle input) that they say will be sure to help you win. "The first column is what your opponent is going to play: A for Rock, B for Paper, and C for Scissors. The second column--" Suddenly, the Elf is called away to help with someone's tent.

The second column, you reason, must be what you should play in response: X for Rock, Y for Paper, and Z for Scissors. Winning every time would be suspicious, so the responses must have been carefully chosen.

The winner of the whole tournament is the player with the highest score. Your total score is the sum of your scores for each round. The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

Since you can't be sure if the Elf is trying to help you or trick you, you should calculate the score you would get if you were to follow the strategy guide.

For example, suppose you were given the following strategy guide:

A Y
B X
C Z
This strategy guide predicts and recommends the following:

In the first round, your opponent will choose Rock (A), and you should choose Paper (Y). This ends in a win for you with a score of 8 (2 because you chose Paper + 6 because you won).
In the second round, your opponent will choose Paper (B), and you should choose Rock (X). This ends in a loss for you with a score of 1 (1 + 0).
The third round is a draw with both players choosing Scissors, giving you a score of 3 + 3 = 6.
In this example, if you were to follow the strategy guide, you would get a total score of 15 (8 + 1 + 6).

What would your total score be if everything goes exactly according to your strategy guide?


Opponent
A = Rock
B = Paper
C = Scissors

Me
X = Rock
Y = Paper
Z = Scissors


Rock = 1p
Paper = 2p
Scissors = 3p

Win = 6p
Loss = 0p
Draw = 3p

Rock beats Scissors
Paper beats Rock
Scissors beats Rock

	X	Y	Z
A	D	W	L
B	L	D	W
C	W	L	D

*/

package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"AoC2022/aoc_helpers"
)

func decodeMove(m string) int {
	var ret int = -1

	switch {
	case m == "A":
		ret = 0
	case m == "X":
		ret = 0
	case m == "B":
		ret = 1
	case m == "Y":
		ret = 1
	case m == "C":
		ret = 2
	case m == "Z":
		ret = 2
	default:
		fmt.Println("Unknown input: ", m)
	}
	return ret
}

var partFlag = flag.Int("p", 0, "part 0 (default) or part 1")
var dbgFlag = flag.Bool("d", false, "debug flag")

func main() {
	var a = [3][3]int{{4, 8, 3}, {1, 5, 9}, {7, 2, 6}}

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	if *partFlag > 1 {
		fmt.Println("p flag not 0 or 1, aborting!")
		return
	}

	lines, err := aoc_helpers.ReadLines(args[0])
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	score := 0
	for i, line := range lines {
		if *dbgFlag {
			fmt.Print("round ", i, ": ", line)
		}
		s := strings.Split(line, " ")

		// Add points for win or draw
		score += a[decodeMove(s[0])][decodeMove(s[1])]
		if *dbgFlag {
			fmt.Println("--> score", score)
		}
	}

	fmt.Println("Total score:", score)
}
