package numbers

// Checks if a number is prime or not
func IsPrime(num int) bool {
    for i := 2; i < num; i++ {
        if num%i == 0 {
            return false
        }
	}
    return true
}
