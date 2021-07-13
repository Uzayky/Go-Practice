package Queens_Attack_II

import "math"

// https://www.hackerrank.com/challenges/queens-attack-2/problem

func QueensAttack(boardSize int32, obstacleSize int32, queenRow int32, queenColumn int32, obstacles [][]int32) int32 {

	var totalThreatenedSquares int32
	totalThreatenedSquaresByQueen := make(map[string]int32)
	totalThreatenedSquaresByObstacles := make(map[string]int32)

	CalculateThreatenedSquaresByQueen(boardSize, queenRow, queenColumn, totalThreatenedSquaresByQueen)
	for _, threatenedSquaresByQueenForDirection := range totalThreatenedSquaresByQueen {
		totalThreatenedSquares += threatenedSquaresByQueenForDirection
	}
	for obstacleIndex := 0; int32(obstacleIndex) < obstacleSize; obstacleIndex++ {
		CalculateThreatenedSquaresByObstacle(boardSize, queenRow, queenColumn, obstacles[obstacleIndex][0], obstacles[obstacleIndex][1], totalThreatenedSquaresByObstacles)
	}
	for _, threatenedSquaresByObstacleForDirection := range totalThreatenedSquaresByObstacles {
		totalThreatenedSquares -= threatenedSquaresByObstacleForDirection
	}

	return totalThreatenedSquares
}

func CalculateThreatenedSquaresByQueen(boardSize int32, row int32, column int32, threatenedSquaresByQueen map[string]int32) {

	threatenedSquaresByQueen["right"] = boardSize - column
	threatenedSquaresByQueen["left"] = column - 1
	threatenedSquaresByQueen["top"] = boardSize - row
	threatenedSquaresByQueen["bot"] = row - 1
	threatenedSquaresByQueen["topRight"] = CalculateThreatenedCrossSquares(threatenedSquaresByQueen["right"], threatenedSquaresByQueen["top"])
	threatenedSquaresByQueen["topLeft"] = CalculateThreatenedCrossSquares(threatenedSquaresByQueen["left"], threatenedSquaresByQueen["top"])
	threatenedSquaresByQueen["botRight"] = CalculateThreatenedCrossSquares(threatenedSquaresByQueen["right"], threatenedSquaresByQueen["bot"])
	threatenedSquaresByQueen["botLeft"] = CalculateThreatenedCrossSquares(threatenedSquaresByQueen["left"], threatenedSquaresByQueen["bot"])
}

func CalculateThreatenedCrossSquares(threatenedLinearSquaresOne int32, threatenedLinearSquaresTwo int32) int32 {

	if threatenedLinearSquaresOne > threatenedLinearSquaresTwo {
		return threatenedLinearSquaresTwo
	} else {
		return threatenedLinearSquaresOne
	}
}

func CalculateThreatenedSquaresByObstacle(boardSize int32, queenRow int32, queenColumn int32, obstacleRow int32, obstacleColumn int32, threatenedSquaresByObstacles map[string]int32) {

	absRow := int32(math.Abs(float64(queenRow - obstacleRow)))
	absColumn := int32(math.Abs(float64(queenColumn - obstacleColumn)))

	if obstacleRow == queenRow && obstacleColumn > queenColumn {
		if threatenedSquaresByObstacles["right"] == 0 || boardSize-obstacleColumn+1 > threatenedSquaresByObstacles["right"] {
			threatenedSquaresByObstacles["right"] = boardSize - obstacleColumn + 1
		}
	} else if obstacleRow == queenRow && obstacleColumn < queenColumn {
		if threatenedSquaresByObstacles["left"] == 0 || obstacleColumn > threatenedSquaresByObstacles["left"] {
			threatenedSquaresByObstacles["left"] = obstacleColumn
		}
	} else if obstacleColumn == queenColumn && obstacleRow > queenRow {
		if threatenedSquaresByObstacles["top"] == 0 || boardSize-obstacleRow+1 > threatenedSquaresByObstacles["top"] {
			threatenedSquaresByObstacles["top"] = boardSize - obstacleRow + 1
		}
	} else if obstacleColumn == queenColumn && obstacleRow < queenRow {
		if threatenedSquaresByObstacles["bot"] == 0 || obstacleRow > threatenedSquaresByObstacles["bot"] {
			threatenedSquaresByObstacles["bot"] = obstacleRow
		}
	} else if queenRow > obstacleRow && queenColumn < obstacleColumn && absRow == absColumn {
		if threatenedSquaresByObstacles["botRight"] == 0 || CalculateThreatenedCrossSquares(obstacleRow, boardSize-obstacleColumn+1) > threatenedSquaresByObstacles["botRight"] {
			threatenedSquaresByObstacles["botRight"] = CalculateThreatenedCrossSquares(obstacleRow, boardSize-obstacleColumn+1)
		}
	} else if queenRow < obstacleRow && queenColumn < obstacleColumn && absRow == absColumn {
		if threatenedSquaresByObstacles["topRight"] == 0 || CalculateThreatenedCrossSquares(boardSize-obstacleRow+1, boardSize-obstacleColumn+1) > threatenedSquaresByObstacles["topRight"] {
			threatenedSquaresByObstacles["topRight"] = CalculateThreatenedCrossSquares(boardSize-obstacleRow+1, boardSize-obstacleColumn+1)
		}
	} else if queenRow < obstacleRow && queenColumn > obstacleColumn && absRow == absColumn {
		if threatenedSquaresByObstacles["topLeft"] == 0 || CalculateThreatenedCrossSquares(boardSize-obstacleRow+1, obstacleColumn) > threatenedSquaresByObstacles["topLeft"] {
			threatenedSquaresByObstacles["topLeft"] = CalculateThreatenedCrossSquares(boardSize-obstacleRow+1, obstacleColumn)
		}
	} else if queenRow > obstacleRow && queenColumn > obstacleColumn && absRow == absColumn {
		if threatenedSquaresByObstacles["botLeft"] == 0 || CalculateThreatenedCrossSquares(obstacleRow, obstacleColumn) > threatenedSquaresByObstacles["botLeft"] {
			threatenedSquaresByObstacles["botLeft"] = CalculateThreatenedCrossSquares(obstacleRow, obstacleColumn)
		}
	}
}
