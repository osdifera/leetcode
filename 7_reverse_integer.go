package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {

	var result int
	for x!=0{
        result=result*10+x%10
        if result > math.MaxInt32  || result < math.MinInt32 {
            return 0
        }
        x/=10
    }
	return result
}

func main() {
	fmt.Println("***Executing***")
	result := reverse(12345)
	fmt.Println(result)
}


//10+32mod10=32
