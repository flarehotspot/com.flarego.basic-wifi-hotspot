package controllers

func divideIntoTensFivesOnes(n int) (int, int, int) {
	tens := n / 10
	remainder := n % 10
	fives := remainder / 5
	ones := remainder % 5
	return tens, fives, ones
}
