package grains

import "math"

func Square(board_size float64) []int {
	// start with 0 grains, increase with each tile
	var grain_count float64 = 0
	// create an empty slice (since arrays have a fixed size)
	var game_board []int
	// tile is the instance we are on, increase with each tile
	for tile := 0.0; tile < board_size; tile++ {
		var grains float64 = math.Exp2(tile)
		grain_count += grains
		game_board[int(tile)] = int(grains)
	}
	return game_board
}

func Total(board_size float64) int {
	// start with 0 grains, increase with each tile
	var grain_count float64 = 0
	// create an empty slice (since arrays have a fixed size)
	var game_board []int
	// tile is the instance we are on, increase with each tile
	for tile := 0.0; tile < board_size; tile++ {
		var grains float64 = math.Exp2(tile)
		grain_count += grains
		game_board[int(tile)] = int(grains)
	}
	return int(grain_count)
}
