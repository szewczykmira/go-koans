package go_koans

func isPrimeNumber(possiblePrime int) bool {
	for underPrime := 2; underPrime < possiblePrime; underPrime++ {
		if possiblePrime%underPrime == 0 {
			return false
		}
	}
	return true
}

func findPrimeNumbers(channel chan int) {
	for i := 2; i < 100;  i++ {
		if isPrimeNumber(i){
			channel <- i
		}
	}
}

func aboutConcurrency() {
	ch := make(chan int)

	assert(__delete_me__) // concurrency can be almost trivial
	// your code goes here

	assert(<-ch == 2)
	assert(<-ch == 3)
	assert(<-ch == 5)
	assert(<-ch == 7)
	assert(<-ch == 11)
}
