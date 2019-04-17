package main

type XY struct {
	X  int32
	Y  int32
	HV string
}

// init Queue
var queue = []XY{}

func Horizontal(grid []string, result [][]int32, X int32, Y int32) [][]int32 {
	var count int32 = 1
	//left
	for {
		if Y == 0 {
			break
		}
		if Y-count < 0 || grid[X][Y-count:Y-count+1] == "X" {
			break
		} else {
			if result[X][Y-count] == 0 && result[X][Y-count] < result[X][Y]+1 {
				xy := XY{X, Y - count, "horizontal"}
				queue = append(queue, xy)
				result[X][Y-count] = result[X][Y] + 1
			}
			count++
		}
	}
	count = 1
	//right
	for {
		if Y+count == int32(len(grid)) {
			break
		}
		if Y+count > int32(len(grid)) || grid[X][Y+count:Y+count+1] == "X" {
			break
		} else {
			if result[X][Y+count] == 0 && result[X][Y+count] < result[X][Y]+1 {
				xy := XY{X, Y + count, "horizontal"}
				queue = append(queue, xy)
				result[X][Y+count] = result[X][Y] + 1
			}
			count++
		}
	}
	return result
}

func Veritical(grid []string, result [][]int32, X int32, Y int32) [][]int32 {
	var count int32 = 1

	//top
	for {
		if X == 0 || X-count < 0 {
			break
		}
		if X-count < 0 || grid[X-count][Y:Y+1] == "X" {
			break
		} else {
			if result[X-count][Y] == 0 && result[X-count][Y] < result[X][Y]+1 {
				xy := XY{X - count, Y, "Veritical"}
				queue = append(queue, xy)
				result[X-count][Y] = result[X][Y] + 1
			}
			count++
		}
	}
	count = 1
	//down
	for {
		if X+count == int32(len(grid)) {
			break
		}
		if X+count > int32(len(grid)) || grid[X+count][Y:Y+1] == "X" {
			break
		} else {
			if result[X+count][Y] == 0 && result[X+count][Y] < result[X][Y]+1 {
				result[X+count][Y] = result[X][Y] + 1
				xy := XY{X + count, Y, "Veritical"}
				queue = append(queue, xy)
			}
			count++
		}
	}
	return result
}

// Complete the minimumMoves function below.
func minimumMoves(grid []string, startX int32, startY int32, goalX int32, goalY int32) int32 {
	// Make 2D slice
	result := make([][]int32, len(grid))
	for i := 0; i < len(grid); i++ {
		result[i] = make([]int32, len(grid))
	}

	result[startX][startY] = 0
	xy := XY{startX, startY, "horizontal"}
	queue = append(queue, xy)
	xy2 := XY{startX, startY, "Veritical"}
	queue = append(queue, xy2)

	for {
		if len(queue) == 0 {
			break
		} else if len(queue) > 0 && queue[0].HV == "horizontal" {
			result = Veritical(grid, result, queue[0].X, queue[0].Y)
			queue = queue[1:]
		} else if len(queue) > 0 && queue[0].HV == "Veritical" {
			result = Horizontal(grid, result, queue[0].X, queue[0].Y)
			queue = queue[1:]
		}
	}

	return result[goalX][goalY]
}
