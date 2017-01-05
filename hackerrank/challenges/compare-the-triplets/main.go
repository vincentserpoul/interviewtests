package main

import (
	"fmt"
	"log"
)

func main() {
	challenger1scores, challenger2scores, err := getInputs()
	if err != nil {
		log.Fatal(err)
	}

	compareScore1, compareScore2 := compareScores(challenger1scores, challenger2scores)

	fmt.Printf("%d %d\n", compareScore1, compareScore2)
}

// Score represent the score of a challenger
type Score struct {
	Zero int
	One  int
	Two  int
}

func getInputs() (challenger1scores Score, challenger2scores Score, err error) {

	_, err = fmt.Scanf(
		"%d %d %d\n%d %d %d",
		&challenger1scores.Zero, &challenger1scores.One, &challenger1scores.Two,
		&challenger2scores.Zero, &challenger2scores.One, &challenger2scores.Two,
	)

	return challenger1scores, challenger2scores, err
}

func compareScores(challenger1scores Score, challenger2scores Score) (compareScore1 int, compareScore2 int) {

	if challenger1scores.Zero > challenger2scores.Zero {
		compareScore1++
	} else if challenger1scores.Zero < challenger2scores.Zero {
		compareScore2++
	}

	if challenger1scores.One > challenger2scores.One {
		compareScore1++
	} else if challenger1scores.One < challenger2scores.One {
		compareScore2++
	}

	if challenger1scores.Two > challenger2scores.Two {
		compareScore1++
	} else if challenger1scores.Two < challenger2scores.Two {
		compareScore2++
	}

	return compareScore1, compareScore2
}
