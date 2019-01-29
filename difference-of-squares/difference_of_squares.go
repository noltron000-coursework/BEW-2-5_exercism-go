package diffsquares

func SumOfSquares(number int) (int) {
	var sum int = 0
	for n := 0; n<=number; n++ {
		sum += n * n
	}
	return sum
}

func SquareOfSum(number int) (int) {
	var sum int = 0
	for n := 0; n<=number; n++ {
		sum += n
	}
	return sum * sum
}

func Difference(number int) (int) {
	var diffy int = 0
	diffy = SquareOfSum(number) - SumOfSquares(number)
	return diffy
}