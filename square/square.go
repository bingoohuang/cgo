package square

/* int square(int x) {
  return x * x;
}

  int sumFirstNSquares(int n) {
		int total = 0;
		for (int i = 0; i < n; i++) {
      total += square(i);
		}
		return total;
	}
*/
import "C"

func sumFirstNSquares(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += int(C.square(C.int(i)))
	}
	return total
}

func betterSumFirstNSquares(n int) int {
	return int(C.sumFirstNSquares(C.int(n)))
}

func goSumFirstNSquares(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += i * i
	}
	return total
}
