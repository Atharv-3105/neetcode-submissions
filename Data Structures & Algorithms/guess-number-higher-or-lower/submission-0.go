/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    left := 1
	right := n
	for left <= right {
		num := left + (right - left) / 2
		response := guess(num)
		if response == 0 {
			return num
		}else if response == -1 {
			right = num - 1
		}else {
			left = num + 1
		}
	}
	return -1
}
