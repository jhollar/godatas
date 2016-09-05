package asymptotic

import (
	"fmt"
)

// Asymptotic Analysisto
//
// Why is it called Asymptotic Analysis?
//
// In every case for a given function f(n) we are trying to find another function g(n) which approximates f(n) at higher values of n.
// That means g(n) is also a curve which approximates f(n) at higher values of n.
//
// In mathematics we call such a curve an asymptotic curve. In other terms, g(n) is the asymptotic curve for f(n).
// For this reason, we call algorithm analysis asymptotic analysis

// Loops - The running time of a loop is, at most, the running time of the statements inside the loop (including tests) multiplied
//         by the number of iterations
//
// Total time = a constant C x N = CN = 0(n)

func Loops(n int) {

	for i := 0; i < n; i++ {

	}
	fmt.Printf("looped %d times\n", n)
}

// NestedLoop - Analyze from the inside out. Total running time is the product of the sizes of all the loops.
//
// Total time = C x N x N = CN^2 = 0(n^2)
//
func NestedLoop(n int) {

	// outer loop executed n times
	for i := 0; i < n; i++ {
		// inner loop executes n times
		for j := 0; j < n; j++ {
			fmt.Printf("NestedLoop - i value %d and j value %d\n", i, j)
		}
	}
}

// Consecutive statements: Add the time complexities of each statement.
//
// Total Time = O(n^2)
//
func ConsecutiveStatements(n int) {

	// executes n times
	for i := 0; i < n; i++ {
		fmt.Printf("ConsecutiveStatements - %d\n", i)
	}

	// Outer loop executed n times
	for i := 0; i < n; i++ {
		// inner loop executes n times
		for j := 0; j < n; j++ {
			fmt.Printf("ConsecutiveStatements - i value %d and j value %d\n", i, j)
		}
	}
}

// If-then-else statements: Worst-case running time: the test, plus either the then part or the else part
// (whichever is the larger)
//
// Total Time = O(n)
//
func IfThenElseStatements(n int) {

	if n == 1 {
		fmt.Printf("Wrong Value - %d\n", n)
	} else {
		for i := 0; i < n; i++ {
			fmt.Printf("IfThenElseStatements - %d\n", i)
		}
	}

}

// Logarithmic complexity: An algorithm is O(logn) if it takes a constant value to cut the problem size by a fraction
//       (usually by 1/2). Aas an example condider the following

// Total Time = O(logn)
//
func LogarithmicComplexity(n int) {

	i := 1
	for i <= n {
		i = i * 2
		fmt.Printf("LogarithmicComplexity -  %d\n", i)
	}

}

// The worst case rate of growth is O(logn). The same discussion holds good for the decreasing sequence as well

// Total Time = O(logn)
//
func LogarithmicComplexityWorst(n int) {

	i := n
	for i >= 1 {
		i = i / 2
		fmt.Printf("LogarithmicComplexityWorst -  %d\n", i)
	}

}

// What is the complexity of this function:   0(n^2 logn)

func LogarithmicComplexityLogn1(n int) {

	count := 0

	for i := (n / 2); i < n; i++ {
		j := 1

		for (j + n/2) <= n {
			k := 1
			for k <= n {
				count = count + 1
				k = k * 2
			}
			j = j + 1
		}
	}

	fmt.Printf("LogarithmicComplexityLogn1 -  %d\n", count)

}

// What is the complexity of this function:   0(logn^2)

func LogarithmicComplexityLogn2(n int) {

	count := 0

	for i := (n / 2); i < n; i++ {
		j := 1

		for (j + n/2) <= n {
			k := 1
			for k <= n {
				count = count + 1
				k = k * 2
			}
			j = j * 2
		}
	}

	fmt.Printf("LogarithmicComplexityLogn2 -  %d\n", count)

}

// What is the complexity of this recursive function. The recurrence for this code is clearly T(n) = T(n -3) + cn^2 for
// some constant c > 0 since each call prints out n^2 asterisks and calls itself recursively on n - 3. Using the iterative method
// we get T(n) = T(n - 3) + cn^2.
//
// Using the Subtraction and Conquer master theorem, we get T(n) = Theta(n^3)
//

func RecursiveFunction1(n int) {

	count := 0
	if n <= 0 {
		return
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			count = count + 1
		}
	}

	RecursiveFunction1(n - 3)
	fmt.Printf("RecursiveFunction1 -  %d\n", count)

}
