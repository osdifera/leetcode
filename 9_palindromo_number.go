package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	var reverse, origin int
	origin = x

	for x!=0{
		reverse = reverse*10+x%10
		x = x/10
	}

	return origin==reverse
}

func main09() {
	result := isPalindrome(4199141)
	fmt.Println(result)
}