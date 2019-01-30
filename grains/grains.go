package grains

import (
	"errors"
	"math"
)

func Square(input int) (uint64, error) {
	// initialize without any errors
	var err error = nil
	// check for error in input
	if input < 1 || input > 64 {
		err = errors.New("input must be ≤ 1")
	}

	// get input, set to uint64
	var num_tiles uint64 = uint64(input)
	// start with 0 grains
	var tile_grains uint64 = 0

	// sets 2^tile power to tile_grains, the # grains on the tile
	tile_grains += uint64(math.Exp2(float64(num_tiles - 1)))

	return tile_grains, err
}

func Total() uint64 {
	// no input, set num_tiles to 64
	var num_tiles uint64 = 64
	// start with 0 grains
	var total_grains uint64 = 0

	for tile := uint64(0); tile < num_tiles; tile++ {
		// adds 2^tile power to total_grains, the # grains on the board
		total_grains += uint64(math.Exp2(float64(tile)))
	}
	return total_grains
}
