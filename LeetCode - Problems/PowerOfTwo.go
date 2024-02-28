/*
Given an integer n, return true if it is a power of two. Otherwise, return false.
An integer n is a power of two, if there exists an integer x such that n == 2x.
*/

package PowerOfTwo

func isPowerOfTwo(n int) bool {

	if n%2 == 1 && n != 1 {
		return false
	}

	for i := 0; i < 31; i++ {
		if n == 1<<i {
			return true
		}
	}
	return false
}
