package main

import (
	"fmt"
	"math"
)

type Coords struct {
	X int
	Y int
}

func main() {
	boardSize, queenPos, obstaclesPos := getInputArray()
	fmt.Printf("%d\n", GetAvailBoxes(boardSize, queenPos, obstaclesPos))
}

// boardsize obstaclesCount
// QueenPos.X QueenPos.Y
// Obstacle.X Obstacle.Y...
func getInputArray() (boardSize int, queenPos Coords, obstaclesPos []Coords) {
	obstaclesCount := 0
	var x, y int
	fmt.Scanf("%d %d\n%d %d\n", &boardSize, &obstaclesCount, &y, &x)

	queenPos.X = x - 1
	queenPos.Y = y - 1

	for i := 0; i < obstaclesCount; i++ {
		var x, y int
		fmt.Scanf("%d %d", &y, &x)
		obstaclesPos = append(obstaclesPos, Coords{Y: y - 1, X: x - 1})
	}

	return boardSize, queenPos, obstaclesPos
}

func GetAvailBoxes(boardSize int, queenPos Coords, obstacles []Coords) (totalAvailBoxes int) {

	top := Coords{X: queenPos.X, Y: boardSize - 1}
	topRight := Coords{X: queenPos.X + int(math.Min(float64(boardSize-1-queenPos.Y), float64(boardSize-1-queenPos.X))), Y: boardSize - 1}
	right := Coords{X: boardSize - 1, Y: queenPos.Y}
	bottomRight := Coords{X: queenPos.X + int(math.Min(float64(queenPos.Y), float64(boardSize-1-queenPos.X))), Y: queenPos.Y - int(math.Min(float64(queenPos.Y), float64(boardSize-1-queenPos.X)))}
	bottom := Coords{X: queenPos.X, Y: 0}
	bottomLeft := Coords{X: queenPos.X - int(math.Min(float64(queenPos.Y), float64(queenPos.X))), Y: queenPos.Y - int(math.Min(float64(queenPos.Y), float64(queenPos.X)))}
	left := Coords{X: 0, Y: queenPos.Y}
	topLeft := Coords{X: queenPos.X - int(math.Min(float64(boardSize-1-queenPos.Y), float64(queenPos.X))), Y: boardSize - 1}
	extremes := []Coords{top, topRight, right, bottomRight, bottom, bottomLeft, left, topLeft}

	// For each of the 8 directions from the queen
	for _, extremePos := range extremes {
		normQEVect := int((math.Max(
			math.Abs(float64(extremePos.X-queenPos.X)),
			math.Abs(float64(extremePos.Y-queenPos.Y)),
		)))
		// default number of boxes available (= vector norm)
		currBoxesAvail := normQEVect

		// For each obstacle
		for _, obstaclePos := range obstacles {
			// Is it colinear?
			if (extremePos.X-queenPos.X)*(obstaclePos.Y-queenPos.Y) == (extremePos.Y-queenPos.Y)*(obstaclePos.X-queenPos.X) {
				// Is it in the same direction?
				if (math.Pow(float64((extremePos.X-queenPos.X)+(obstaclePos.X-queenPos.X)), 2) +
					math.Pow(float64((extremePos.Y-queenPos.Y)+(obstaclePos.Y-queenPos.Y)), 2)) >
					(math.Pow(float64((extremePos.X-queenPos.X)), 2) +
						math.Pow(float64((extremePos.Y-queenPos.Y)), 2)) {
					// How many boxes removed (obstacle vector length)?
					obsBoxesAvail := int(math.Max(
						math.Abs(float64(obstaclePos.X-queenPos.X))-1,
						math.Abs(float64(obstaclePos.Y-queenPos.Y))-1,
					))
					if obsBoxesAvail < currBoxesAvail {
						currBoxesAvail = obsBoxesAvail
					}
				}
			}
		}

		totalAvailBoxes += int(currBoxesAvail)
	}

	return totalAvailBoxes
}

// old brute force way
// func boxesAvail(boardSize int, queenPos Coords, obstaclesPos []Coords) (boxesCount int) {

// 	var i, j int

// 	// Check top
// 	i = queenPos.Y + 1
// 	for i <= boardSize-1 {
// 		// fmt.Printf("check top\n")
// 		if isObstacleHere(Coords{Y: i, X: queenPos.X}, obstaclesPos) {
// 			break
// 		}
// 		i++
// 		boxesCount++
// 	}

// 	// Count top avail boxes
// 	topAvailBoxes := boardSize - 1 - queenPos.Y

// 	// Check bot
// 	i = queenPos.Y - 1
// 	for i >= 0 {
// 		// fmt.Printf("check bottom\n")
// 		if isObstacleHere(Coords{Y: i, X: queenPos.X}, obstaclesPos) {
// 			break
// 		}
// 		i--
// 		boxesCount++
// 	}

// 	// Count bot avail boxes
// 	botAvailBoxes := queenPos.Y

// 	// Check left
// 	j = queenPos.X - 1
// 	for j >= 0 {
// 		// fmt.Printf("check left\n")
// 		if isObstacleHere(Coords{Y: queenPos.Y, X: j}, obstaclesPos) {
// 			break
// 		}
// 		j--
// 		boxesCount++
// 	}

// 	// Count left avail boxes
// 	leftAvailBoxes := queenPos.X

// 	// Check right
// 	j = queenPos.X + 1
// 	for j <= boardSize-1 {
// 		// fmt.Printf("check right\n")
// 		if isObstacleHere(Coords{Y: queenPos.Y, X: j}, obstaclesPos) {
// 			break
// 		}
// 		j++
// 		boxesCount++
// 	}

// 	// Count right avail boxes
// 	rightAvailBoxes := boardSize - 1 - queenPos.X

// 	// Check top-left
// 	i = queenPos.Y + 1
// 	j = queenPos.X - 1
// 	for j >= 0 && i <= boardSize-1 {
// 		// fmt.Printf("check top left\n")
// 		if isObstacleHere(Coords{Y: i, X: j}, obstaclesPos) {
// 			break
// 		}
// 		i++
// 		j--
// 		boxesCount++
// 	}

// 	// Count top left avail boxes
// 	topleftAvailBoxes := int(math.Min(float64(leftAvailBoxes), float64(topAvailBoxes)))

// 	// Check top-right
// 	i = queenPos.Y + 1
// 	j = queenPos.X + 1
// 	for j <= boardSize-1 && i <= boardSize-1 {
// 		// fmt.Printf("check top right\n")
// 		if isObstacleHere(Coords{Y: i, X: j}, obstaclesPos) {
// 			break
// 		}
// 		i++
// 		j++
// 		boxesCount++
// 	}

// 	// Count top right avail boxes
// 	toprightAvailBoxes := int(math.Min(float64(rightAvailBoxes), float64(topAvailBoxes)))

// 	// Check bottom-left
// 	i = queenPos.Y - 1
// 	j = queenPos.X - 1
// 	for j >= 0 && i >= 0 {
// 		// fmt.Printf("check bottom left\n")
// 		if isObstacleHere(Coords{Y: i, X: j}, obstaclesPos) {
// 			break
// 		}
// 		i--
// 		j--
// 		boxesCount++
// 	}

// 	// Count bottom left avail boxes
// 	bottomleftAvailBoxes := int(math.Min(float64(leftAvailBoxes), float64(botAvailBoxes)))

// 	// Check bottom-right
// 	i = queenPos.Y - 1
// 	j = queenPos.X + 1
// 	for j <= boardSize-1 && i >= 0 {
// 		// fmt.Printf("check bottom right\n")
// 		if isObstacleHere(Coords{Y: i, X: j}, obstaclesPos) {
// 			break
// 		}
// 		i--
// 		j++
// 		boxesCount++
// 	}

// 	// Count bottom right avail boxes
// 	bottomrightAvailBoxes := int(math.Min(float64(rightAvailBoxes), float64(botAvailBoxes)))

// 	return topAvailBoxes + topleftAvailBoxes + leftAvailBoxes + bottomleftAvailBoxes + botAvailBoxes + bottomrightAvailBoxes + rightAvailBoxes + toprightAvailBoxes
// }

// func isObstacleHere(toCheck Coords, obstacles []Coords) bool {
// 	for _, obst := range obstacles {
// 		if toCheck == obst {
// 			return true
// 		}
// 	}
// 	return false
// }
